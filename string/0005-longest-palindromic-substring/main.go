/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/3 下午6:11
***********************************************/

package main

import "fmt"

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
// Related Topics 字符串 动态规划
func main() {
	fmt.Println(longestPalindrome2("babad"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i + l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l + 1 > len(ans) {
				ans = s[i:i+l+1]
			}
		}
	}
	return ans
}

func longestPalindrome2(s string) string {
	// 动态规划
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	// dp[i][i] = true
	// dp[i][i+1] = (s[i] == s[i+1])
	n, ans := len(s), ""
	dp := make([][]int, n)
	for i := 0; i < n; i ++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j+i < n; j ++ {
			k := j + i
			if i == 0 {
				dp[j][k] = 1
			} else if i == 1 {
				if s[j] == s[k] {
					dp[j][k] = 1
				}
			} else {
				if s[j] == s[k] {
					dp[j][k] = dp[j+1][k-1]
				}
			}
			if dp[j][k] > 0 && i + 1 > len(ans) {
				ans = s[j:j+i+1]
			}
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

