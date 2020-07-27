/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/25 下午11:33
***********************************************/

package main

import (
	"fmt"
	"math"
)

//给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。
//
// 注意:
//数组长度 n 满足以下条件:
//
//
// 1 ≤ n ≤ 1000
// 1 ≤ m ≤ min(50, n)
//
//
// 示例:
//
//
//输入:
//nums = [7,2,5,10,8]
//m = 2
//
//输出:
//18
//
//解释:
//一共有四种方法将nums分割为2个子数组。
//其中最好的方式是将其分为[7,2,5] 和 [10,8]，
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
//
// Related Topics 二分查找 动态规划
func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 令 f[i][j]表示将数组的前 i 个数分割为 j 段所能得到的最大连续子数组和的最小值
func splitArray(nums []int, m int) int {
	n := len(nums)
	f := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, m+1)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return f[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
