/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/8 下午10:52
***********************************************/

package main

import "fmt"

//在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//
// 示例:
//
// 输入:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//
//输出: 4
// Related Topics 动态规划
func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 公式：dp[i][j] = min(dp[i-1][j], dp[i][j-], dp[i-1][j-1]) + 1
func maximalSquare(matrix [][]byte) int {
	max := 0
	dp := make([][]int, len(matrix))
	// 1. 先初始化
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				max = 1
			}
		}
	}
	fmt.Printf("%#v\n", dp)
	// 2. 然后
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	fmt.Printf("%#v\n", dp)
	return max * max
}

//leetcode submit region end(Prohibit modification and deletion)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
