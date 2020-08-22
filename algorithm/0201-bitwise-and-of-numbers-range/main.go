/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/23 上午12:19
***********************************************/

package main

// Brian Kernighan's algorithm
func rangeBitwiseAnd(m, n int) int {
	for n > m {
		n &= n-1
	}
	return n
}

// 位移
func rangeBitwiseAnd3(m, n int) int {
	if m == 0 {
		return 0
	}
	moved := 0
	for m < n {
		m, n = m >> 1, n >> 1
		moved ++
	}
	return m << moved
}

// 超时
func rangeBitwiseAnd2(m, n int) int {
	if m == n {
		return m & n
	}
	res := n
	for m < n {
		res = res & m
		m ++
	}
	return res
}
