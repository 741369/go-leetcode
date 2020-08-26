/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午8:23
***********************************************/

package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}

	var req float64 // 请求值
	req = 3

	var resp *float64
	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f * %f = %f\n", req, req, *resp)
}
