/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/27 下午10:13
***********************************************/

package main

import "fmt"

//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//
//
//
// 示例 1:
//
// 输入: [1,2,0]
//输出: 3
//
//
// 示例 2:
//
// 输入: [3,4,-1,1]
//输出: 2
//
//
// 示例 3:
//
// 输入: [7,8,9,11,12]
//输出: 1
//
//
//
//
// 提示：
//
// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
// Related Topics 数组
func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

// 数组交换
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

//leetcode submit region begin(Prohibit modification and deletion)
// 默认数组实现哈希表
func firstMissingPositive2(nums []int) int {
	n := len(nums)

	// 将所有小于0的值都赋值为n，
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 将所有1-n的位置都设置为负数
	for i := 0; i < n; i++ {
		tmp := abs(nums[i])
		if tmp < n+1 {
			nums[tmp-1] = -abs(nums[tmp-1])
		}
	}

	// 第一个整数的位置+1即为最小正整数，或者n+1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
