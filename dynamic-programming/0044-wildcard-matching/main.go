/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/5 下午10:40
***********************************************/

package main
//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//
//
// 两个字符串完全匹配才算匹配成功。
//
// 说明:
//
//
// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
//
//
// 示例 1:
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//
// 示例 2:
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
//
//
// 示例 3:
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
//
//
// 示例 4:
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
//
//
// 示例 5:
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false
// Related Topics 贪心算法 字符串 动态规划 回溯算法
func main() {
	println(isMatch("aa", "a"))
	println(isMatch("aa", "*"))
	println(isMatch("cb", "?a"))
	println(isMatch("adceb", "*a*b"))
	println(isMatch("acdcb", "a*c?b"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	// dp[i][j] 表示字符串s的前i个字符与字符串p的前j个字符是否匹配
	// 1. 当j指定的字符为字符时，dp[i][j] = dp[i-1][j-1] && s[i] == p[j]
	// 2. 当j指定的字符为问号?时，dp[i][j] = dp[i-1][j-1]
	// 1,2 合并为：dp[i][j] = dp[i-1][j-1] 当且紧当 s[i] == p[j] 或者p[j] = ?
	// 3. 当j指定的字符为*时，dp[i][j] = dp[i][j-1] || dp[i-1][j]

	// 默认值，dp[0][0] = true, dp[i][0] = false, dp[0][j] 当j指定的元素前面都为*时，为true，其他都为false

	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	// 初始化存储，避免多次比较
	for i := 0; i <= m; i ++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j ++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i ++ {
		for j := 1; j <= n; j ++ {
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
//leetcode submit region end(Prohibit modification and deletion)


