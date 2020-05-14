/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/14 上午9:55
***********************************************/

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//client := http.DefaultClient
	client := http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

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
