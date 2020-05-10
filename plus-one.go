/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/9 下午4:23
***********************************************/

package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	temp := make([]int, len(digits)+1)
	temp[0] = 1
	return temp
}
