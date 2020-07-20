/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/20 上午8:56
***********************************************/

package main

import "fmt"

//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
// 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
// 说明:
//
//
// 返回的下标值（index1 和 index2）不是从零开始的。
// 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//
//
// 示例:
//
// 输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
// Related Topics 数组 双指针 二分查找
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 双指针 O(n)
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// 二分查找 O(nlogn)
func twoSum(numbers []int, target int) []int {
	right := len(numbers) - 1
	for i := 0; i <= right; i++ {
		left := i + 1
		for left <= right {
			mid := (left + right) / 2
			sum := numbers[i] + numbers[mid]
			if sum == target {
				return []int{i + 1, mid + 1}
			} else if sum > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return nil
}

// 数组遍历
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
