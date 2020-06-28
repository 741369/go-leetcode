/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/28 下午11:16
***********************************************/

package main

import "fmt"

//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回
// 0。
//
//
//
// 示例：
//
// 输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的连续子数组。
//
//
//
//
// 进阶：
//
//
// 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
//
// Related Topics 数组 双指针 二分查找
func main() {
	fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3}))
	fmt.Println(minSubArrayLen(4, []int{1,4,4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(s int, nums []int) int {
	ans := len(nums) + 1
	sum, start, end := 0, 0, 0
	for end < len(nums) {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start ++
		}
		end++
	}


	if ans == len(nums) + 1{
		return 0
	}
	return ans
}
// 暴力法
func minSubArrayLen2(s int, nums []int) int {
	ans := len(nums) + 1
	for i := 0; i < len(nums); i ++ {
		tmp := 0
		for j := i; j < len(nums); j ++ {
			tmp += nums[j]
			if tmp >= s {
				ans = min(j-i+1, ans)
				break
			}
		}
	}
	if ans == len(nums) + 1 {
		return 0
	}
	return ans
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
//leetcode submit region end(Prohibit modification and deletion)

