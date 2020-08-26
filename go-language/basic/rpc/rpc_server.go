/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午5:01
***********************************************/

package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

// MathUtil 用于数学计算
type MathUtil struct {

}

// CalculateCircleArea 计算圆的面积
func (m *MathUtil) CalculateCircleArea(req float64, resp *float64) error {
	*resp = math.Pi * req * req
	return nil
}

func main() {
	// 创建计算实例
	mathUtil := new(MathUtil)

	// 将对象注册到rpc服务中
	err := rpc.Register(mathUtil)
	if err != nil {
		log.Panic(err)
		return
	}

	// 通过该函数把mathUtil中提供的服务注册到HTTP协议上
	// 方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	// 在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
		return
	}
	http.Serve(listen, nil)
}