/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/14 下午11:47
***********************************************/

package main

import (
	"fmt"
	"math"
)

//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//
//
// 例如，给定三角形：
//
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//
//
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//
//
// 说明：
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
// Related Topics 数组 动态规划
func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, 2)

	for i := 0; i < 2 ; i ++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		cur := i % 2
		pre := 1 - cur
		dp[cur][0] = dp[pre][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[cur][j] = min(dp[pre][j-1], dp[pre][j])+triangle[i][j]
		}
		dp[cur][i] = dp[pre][i-1] + triangle[i][i]
	}
	fmt.Printf("%#v\n", dp)
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		res = min(res, dp[(n-1)%2][i])
	}
	return res
}

func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)

	for i := 0; i < n ; i ++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j])+triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	fmt.Printf("%#v\n", dp)
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}