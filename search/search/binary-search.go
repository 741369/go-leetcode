/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/24 下午4:46
***********************************************/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Println(binarySearch(nums, 5))
}

func binarySearch(nums []int, key int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == key {
			return middle
		} else if nums[middle] < key {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	fmt.Println("===========", left, right)
	return -1
}
