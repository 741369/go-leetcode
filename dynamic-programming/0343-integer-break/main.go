/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/30 下午10:04
***********************************************/

package main

import "fmt"

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
//
// 示例 1:
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//
// 示例 2:
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
//
// 说明: 你可以假设 n 不小于 2 且不大于 58。
// Related Topics 数学 动态规划
func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}

//leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
	// dp[i] = max(j * (i-j), j * dp[i-j])
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i ++ {
		tmp := 0
		for j := 1; j < i; j ++ {
			tmp = max(tmp, max(j * (i - j), j * dp[i-j]))
		}
		dp[i] = tmp
	}
	return dp[n]
}
func max(i, j int) int {
	if i > j {
		 return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)

