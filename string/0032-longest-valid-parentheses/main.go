/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/4 上午8:38
***********************************************/

package main
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
// 示例 1:
//
// 输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//
//
// 示例 2:
//
// 输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
//
// Related Topics 字符串 动态规划


//leetcode submit region begin(Prohibit modification and deletion)
// 动态规划方案
func longestValidParentheses(s string) int {
	// 1. () dp[i] = dp[i-2] + 2
	// 2. )) dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
	var (
		ans = 0
		dp = make([]int, len(s))
	)
	for i := 1; i < len(s); i ++ {
		if s[i:i+1] == ")" {
			// 形如 ()
			if s[i-1:i] == "(" {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			// 形如 (())
			} else if i - dp[i-1] > 0 && s[i-dp[i-1]-1:i-dp[i-1]] == "(" {
				if i - dp[i-1] - 2 >= 0 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] +2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
