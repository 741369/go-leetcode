/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/4/28 下午11:56
***********************************************/

package main

import "fmt"

func main() {
	fmt.Println("===", singleNumbers([]int{1, 1, 2, 2, 3, 4, 5, 5, 4, 6, 3, 7}))
}

func singleNumbers(nums []int) []int {
	var (
		ret   = 0
		index = 1
	)
	for _, v := range nums {
		ret ^= v
	}

	for ret&index == 0 {
		index <<= 1
	}

	fmt.Println("===", index)
	res := make([]int, 2)
	for _, v := range nums {
		if v&index == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}

	return res
}
