/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/27 下午2:53
***********************************************/

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋
//装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
// 示例 1:
//
// 输入: [2,3,2]
//输出: 3
//解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2:
//
// 输入: [1,2,3,1]
//输出: 4
//解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
// Related Topics 动态规划

package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	/**
	1. 不偷第一家的情况 nums[1:]
	2. 不偷最后一家的情况 nums[:n-1]
	*/
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(myRob(nums[:len(nums)-1]), myRob(nums[1:]))
}
func myRob(nums []int) int {
	pre2, pre1 := 0, 0
	for i := 0; i < len(nums); i++ {
		cur := max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}

//leetcode submit region end(Prohibit modification and deletion)
func max(first int, second int) int {
	if first >= second {
		return first
	} else {
		return second
	}
}
