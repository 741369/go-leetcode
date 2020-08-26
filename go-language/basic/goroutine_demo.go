/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午5:01
***********************************************/

package basic

import "fmt"

func GoroutineDemo() {
	// 快速生成一个1-100的数组，并且计算他们的平方
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0-100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("ch1: >>>>", i)
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <- ch1 // 通道关闭后取值ok=false
			if !ok {
				break
			}
			fmt.Println("ch2: >>>>", i)
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	// 通道关闭后会退出for range循环
	for i := range ch2 {
		fmt.Println(i)
	}
//ch1: >>>> 0
//ch1: >>>> 1
//ch2: >>>> 0
//ch2: >>>> 1
//ch1: >>>> 2
//	0
//	1
//ch2: >>>> 2
//	4
//ch1: >>>> 3
//ch1: >>>> 4
//ch2: >>>> 3
//ch2: >>>> 4
//ch1: >>>> 5
//	9
//	16
}