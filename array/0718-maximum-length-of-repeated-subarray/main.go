/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/1 上午9:15
***********************************************/

package main

import (
	"fmt"
	"reflect"
)

//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
//
// 示例 1:
//
//
//输入:
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出: 3
//解释:
//长度最长的公共子数组是 [3, 2, 1]。
//
//
// 说明:
//
//
// 1 <= len(A), len(B) <= 1000
// 0 <= A[i], B[i] < 100
//
// Related Topics 数组 哈希表 二分查找 动态规划
func main() {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 动态规划
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i ++{
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n-1; i >= 0; i--{
		for j := m-1; j >= 0; j--{
			if A[i] == B[j] {
				// 因为 A[i] == B[j] ，如果A，B的最长公告前缀为A[i+1:] B[j+1:]的子串+1
				dp[i][j] = dp[i+1][j+1]+1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
// 双指针

// 暴力法
func findLength3(A []int, B []int) int {
	ans := 0
	for i := 0; i < len(A); i ++ {
		for j := 0; j < len(B); j ++ {
			k := 0
			for i+k < len(A) && j+k < len(B) && reflect.DeepEqual(A[i:i+k], B[j:j+k]) {
				k ++
			}
			if k > ans {
				ans = k
			}
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

