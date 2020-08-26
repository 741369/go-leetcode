/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午5:01
***********************************************/

package main

import (
	"fmt"
	"net/http"
	"time"
)

// go语言中自带http包，通过调用 http.ListenAndServe(":8000", nil) 方法
// 就能够处理端口监听、请求解析、路由分配、中间件、响应处理，整体思路类似Koa

func main() {
	http.Handle("/", loggingHandler(middlewareHandler(http.HandlerFunc(HelloWorldHandler))))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("服务器错误")
	}
}

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 执行handler之前的逻辑
		fmt.Println("=================before=============")
		next.ServeHTTP(w, r)
		// 执行handler之后的逻辑
		fmt.Println("=================after==============")
	})
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Started %s %s", r.Method, r.PostForm)
		next.ServeHTTP(w, r)
		fmt.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Fprintf(w, "Hello World")
}