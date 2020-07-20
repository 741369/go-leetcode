/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/19 下午11:27
***********************************************/

package main

import "fmt"

//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//
// 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 lef
//t 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
//
// 求所能获得硬币的最大数量。
//
// 说明:
//
//
// 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
//
//
// 示例:
//
// 输入: [3,1,5,8]
//输出: 167
//解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
//
// Related Topics 分治算法 动态规划
func main() {
	fmt.Println("===========", maxCoins([]int{3, 1, 5, 8}))
}

func maxCoins3(nums []int) int {
	n := len(nums)
	rec := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rec[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][n+1]
}

func maxCoins(nums []int) int {
	n := len(nums)

	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}

	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins2(nums []int) int {
	n := len(nums)
	// 初始化
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1

	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return resolve(val, 0, n+1, dp)
}

func resolve(val []int, left, right int, dp [][]int) int {
	if left >= right-1 {
		return 0
	}

	if dp[left][right] != -1 {
		return dp[left][right]
	}

	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		sum += resolve(val, left, i, dp) + resolve(val, i, right, dp)
		dp[left][right] = max(dp[left][right], sum)
	}
	return dp[left][right]
}

//leetcode submit region end(Prohibit modification and deletion)
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
