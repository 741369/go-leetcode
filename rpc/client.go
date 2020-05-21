/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/21 下午6:11
***********************************************/

package main

import (
	"encoding/gob"
	"errors"
	"leetcode/rpc/utils"
	"log"
	"net"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

type ResponseQueryUser struct {
	User
	Msg string
}

func main() {
	gob.Register(ResponseQueryUser{})

	addr := "0.0.0.0:3333"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("dial error: %v\n", err)
	}
	cli := NewClient(conn)

	var correctQuery func(int) (ResponseQueryUser, error)
	var wrongQuery func(int) (ResponseQueryUser, error)

	cli.Call("queryUser", &correctQuery)
	u, err := correctQuery(1)
	if err != nil {
		log.Printf("query error: %v\n", err)
	} else {
		log.Printf("query result: %v %v %v\n", u.Name, u.Age, u.Msg)
	}
	u, err = correctQuery(2)
	if err != nil {
		log.Printf("query error: %v\n", err)
	} else {
		log.Printf("query result: %v %v %v\n", u.Name, u.Age, u.Msg)
	}

	cli.Call("queryUse", &wrongQuery)
	u, err = wrongQuery(1)
	if err != nil {
		log.Printf("query error: %v\n", err)
	} else {
		log.Println(u)
	}

	conn.Close()
}

// Client struct
type Client struct {
	conn net.Conn
}

// NewClient creates a new client
func NewClient(conn net.Conn) *Client {
	return &Client{conn}
}

// Call transforms a function prototype into a function
func (c *Client) Call(name string, fptr interface{}) {
	container := reflect.ValueOf(fptr).Elem()

	f := func(req []reflect.Value) []reflect.Value {
		cliTransport := utils.NewTransport(c.conn)

		errorHandler := func(err error) []reflect.Value {
			outArgs := make([]reflect.Value, container.Type().NumOut())
			for i := 0; i < len(outArgs)-1; i++ {
				outArgs[i] = reflect.Zero(container.Type().Out(i))
			}
			outArgs[len(outArgs)-1] = reflect.ValueOf(&err).Elem()
			return outArgs
		}
		// package request arguments
		inArgs := make([]interface{}, 0, len(req))
		for i := range req {
			inArgs = append(inArgs, req[i].Interface())
		}
		// send request to server
		err := cliTransport.Send(utils.Data{Name: name, Args: inArgs})
		if err != nil { // local network error or encode error
			return errorHandler(err)
		}
		// receive response from server
		rsp, err := cliTransport.Receive()
		if err != nil { // local network error or decode error
			return errorHandler(err)
		}
		if rsp.Err != "" { // remote server error
			return errorHandler(errors.New(rsp.Err))
		}

		if len(rsp.Args) == 0 {
			rsp.Args = make([]interface{}, container.Type().NumOut())
		}
		// unpackage response arguments
		numOut := container.Type().NumOut()
		outArgs := make([]reflect.Value, numOut)
		for i := 0; i < numOut; i++ {
			if i != numOut-1 { // unpackage arguments (except error)
				if rsp.Args[i] == nil { // if argument is nil (gob will ignore "Zero" in transmission), set "Zero" value
					outArgs[i] = reflect.Zero(container.Type().Out(i))
				} else {
					outArgs[i] = reflect.ValueOf(rsp.Args[i])
				}
			} else { // unpackage error argument
				outArgs[i] = reflect.Zero(container.Type().Out(i))
			}
		}

		return outArgs
	}

	container.Set(reflect.MakeFunc(container.Type(), f))
}
