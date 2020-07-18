/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/17 下午11:19
***********************************************/

package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 你可以假设数组中无重复元素。
//
// 示例 1:
//
// 输入: [1,3,5,6], 5
//输出: 2
//
//
// 示例 2:
//
// 输入: [1,3,5,6], 2
//输出: 1
//
//
// 示例 3:
//
// 输入: [1,3,5,6], 7
//输出: 4
//
//
// 示例 4:
//
// 输入: [1,3,5,6], 0
//输出: 0
//
// Related Topics 数组 二分查找
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert2(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// 二分查找
func searchInsert(nums []int, target int) int {
	res := len(nums)
	i, j := 0, res-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] >= target {
			res = mid
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
