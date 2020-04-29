package main

import "fmt"

func main() {
	fmt.Println("==", lastRemaining(5, 3))
}
func lastRemaining2(n, m int) int {
	return find(n, m)
}

func find(n, m int) int {
	if n == 1 {
		return 0
	}
	x := find(n-1, m)
	return (m + x) % n
}

func lastRemaining(n int, m int) int {
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
