/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/14 上午9:55
***********************************************/

package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false, // 表示是否开启http keepalive功能，也即是否重用连接，默认开启(false)
			Proxy:             http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,                             // 表示连接池对所有host的最大链接数量，一般通过压测 二分得该参数近似值，如果太大容易内存吃满
			MaxIdleConnsPerHost: 100,                             // 表示每个host的最大连接数，如果只有1个host，则可以喝MaxIdleConns一样
			IdleConnTimeout:     time.Duration(90) * time.Second, // 空闲timeout设置，也即socket在该时间内没有交互则自动关闭连接，有交互重置为0
		},
		Timeout: 20 * time.Second,
	}
	return client
}

func main() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal("http new request err = ", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request do err = %v", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read resp body err = ", err)
		return
	}
	log.Println(string(bodyBytes))
}
