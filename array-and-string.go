/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/9 下午3:31
***********************************************/

package main

import "fmt"

func main() {
	nums := []int{0, 0, 2, 3}
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	max, secondMax, maxIndex, secondIndex := 0, 0, -1, -1
	for k, v := range nums {
		if v > max {
			max = v
			maxIndex = k
		}
	}
	for k, v := range nums {
		if v > secondMax && v < max {
			secondMax = v
			secondIndex = k
		}
	}
	if max >= secondMax*2 && maxIndex != secondIndex {
		return maxIndex
	}
	return -1
}
