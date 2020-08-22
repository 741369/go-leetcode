/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/22 下午5:50
***********************************************/

package main

import "fmt"

// CSP论文
// Problem: To print in ascending order all primes less than 10000.
// Use an array of processes, SIEVE, in which each process inputs a prime from its predecessor and prints it.
// The process then inputs an ascending stream of numbers from its predecessor and passes them on to its successor,
// suppressing any that are multiples of the original prime.

// 要找出10000以内所有的素数，这里使用的方法是筛法，
// 即从2开始每找到一个素数就标记所有能被该素数整除的所有数。
// 直到没有可标记的数，剩下的就都是素数。
// 下面以找出10以内所有素数为例，借用 CSP 方式解决这个问题。

// Go语言的并发很简单，并且通过提高并发可以提高处理效率
// 协程之间可以通过通信的方式来共享变量
func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 100000000; num ++ {
		origin <- num
	}
	close(origin)
	<- wait
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <- seq
		if ! ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num % prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}