/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午9:59
***********************************************/
//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
// 示例 1:
//
//
//输入:
//"tree"
//
//输出:
//"eert"
//
//解释:
//'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//
//
// 示例 2:
//
//
//输入:
//"cccaaa"
//
//输出:
//"cccaaa"
//
//解释:
//'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//
//
// 示例 3:
//
//
//输入:
//"Aabb"
//
//输出:
//"bbAa"
//
//解释:
//此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
//
// Related Topics 堆 哈希表
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
