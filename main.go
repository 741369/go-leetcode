/**********************************************
** @Des:
** @Author: liuzhiwang
** @Last Modified time: 2020/6/3 下午6:23
***********************************************/

package main

import (
	"fmt"
	"sync"
)

func main() {
	printStr("1234567890")
	fmt.Println("==================done")
}

func printStr(str string) {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		for i := 0; i < len(str); i += 2 {
			<-ch1
			fmt.Println(string(str[i]))
			ch2 <- 1
		}
		<-ch1
		wg.Done()
	}()

	go func() {
		for i := 1; i < len(str); i += 2 {
			<-ch2
			fmt.Println(string(str[i]))
			ch1 <- 1
		}
		wg.Done()
	}()
	ch1 <- 1
	wg.Wait()
}
