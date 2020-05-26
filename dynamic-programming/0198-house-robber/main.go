/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/26 下午10:35
***********************************************/

package main

import "fmt"

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
//如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
// 示例 1:
//
// 输入: [1,2,3,1]
//输出: 4
//解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2:
//
// 输入: [2,7,9,3,1]
//输出: 12
//解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
// Related Topics 动态规划
func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	// 抢了第i个房间就不能抢第i-1个，因此，dp[i] = max(dp[i-2]+nums[i], dp[i-1])
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
