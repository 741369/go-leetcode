/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午9:59
***********************************************/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaaaccabb"
	fmt.Println("==", frequencySort(str))
}

func frequencySort(s string) string {
	strArr := strings.Split(s, "")
	strMap := make(map[string]string)
	for _, val := range strArr {
		strMap[val] += val
	}

	strSlice := make([]string, 0, len(strMap))
	for _, val := range strMap {
		strSlice = append(strSlice, val)
	}
	heapSort(strSlice, len(strSlice))
	return strings.Join(strSlice, "")
}

func heapSort(nums []string, lens int) {
	buildHeap(nums, lens)
	for i := lens - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		lens--
		heap(nums, 0, lens)
	}
}

// 构建最小堆
func buildHeap(nums []string, lens int) {
	for i := lens / 2; i >= 0; i-- {
		heap(nums, i, lens)
	}
}

func heap(nums []string, i int, lens int) {
	left, right, min := 2*i+1, 2*i+2, i
	if left < lens && len(nums[left]) < len(nums[min]) {
		min = left
	}
	if right < lens && len(nums[right]) < len(nums[min]) {
		min = right
	}
	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		heap(nums, min, lens)
	}
}
