/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/21 下午5:15
***********************************************/

package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"leetcode/rpc/utils"
	"log"
	"net"
	"reflect"
)

func main() {
	addr := "0.0.0.0:3333"
	srv := NewServer(addr)

	gob.Register(ResponseQueryUser{})

	srv.Register("queryUser", queryUser)

	// 启动一个服务
	srv.Run()

	select {}
}

func queryUser(uid int) (ResponseQueryUser, error) {
	db := map[int]User{
		0: {
			Name: "lzw",
			Age:  18,
		},
		1: {
			Name: "test",
			Age:  22,
		},
	}
	if data, ok := db[uid]; ok {
		return ResponseQueryUser{
			User: data,
			Msg:  "success",
		}, nil
	}
	return ResponseQueryUser{
		User: User{},
		Msg:  "fail",
	}, errors.New("uid is not in database")
}

type User struct {
	Name string
	Age  int
}

type ResponseQueryUser struct {
	User
	Msg string
}

type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

// 创建一个server
func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

func (srv *Server) Run() {
	l, err := net.Listen("tcp", srv.addr)
	if err != nil {
		log.Fatalf("listen on %s err: %v\n", srv.addr, err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err)
			continue
		}

		go func() {
			srvTransport := utils.NewTransport(conn)
			for {
				// read request from client
				req, err := srvTransport.Receive()
				if err != nil {
					if err != io.EOF {
						log.Printf("read err: %v\n", err)
					}
					return
				}

				// get method by name
				f, ok := srv.funcs[req.Name]
				if !ok { // if method requested does not exist
					e := fmt.Sprintf("func %s does not exist", req.Name)
					log.Println(e)
					if err = srvTransport.Send(utils.Data{
						Name: req.Name,
						Err:  e,
					}); err != nil {
						log.Printf("transport write err: %v\n", err)
					}
					continue
				}
				log.Printf("func %s is called\n", req.Name)
				// unpackage request arguments
				inArgs := make([]reflect.Value, len(req.Args))
				for i := range req.Args {
					inArgs[i] = reflect.ValueOf(req.Args[i])
				}
				// invoke requested method
				out := f.Call(inArgs)
				// package response arguments (except error)
				outArgs := make([]interface{}, len(out)-1)
				for i := 0; i < len(out)-1; i++ {
					outArgs[i] = out[i].Interface()
				}
				// package error argument
				var e string
				if _, ok := out[len(out)-1].Interface().(error); !ok {
					e = ""
				} else {
					e = out[len(out)-1].Interface().(error).Error()
				}
				// send response to client
				err = srvTransport.Send(utils.Data{Name: req.Name, Args: outArgs, Err: e})
				if err != nil {
					log.Printf("transport write err: %v\n", err)
				}
			}
		}()
	}
}

func (srv *Server) Register(name string, f interface{}) {
	if _, ok := srv.funcs[name]; ok {
		return
	}
	srv.funcs[name] = reflect.ValueOf(f)
}
