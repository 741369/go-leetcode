/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/9 下午3:16
***********************************************/

package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("===", pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, v := range nums {
		sum += v
	}
	for k, v := range nums {
		if leftSum == sum-leftSum-v {
			return k
		}
		leftSum += v
	}
	return -1
}
