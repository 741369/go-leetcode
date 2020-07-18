/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/18 下午11:42
***********************************************/

package main

import "fmt"

//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
//
// 示例 1:
//
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出: true
//
//
// 示例 2:
//
// 输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出: false
// Related Topics 字符串 动态规划
func main() {
	fmt.Println("====", isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println("====", isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n+m) != t {
		return false
	}
	bp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		bp[i] = make([]bool, m+1)
	}

	// bp[i][j] = bp[i-1][j] && s3[p] == s1[i-1]
	// 或者 bp[i][j] = bp[i][j-1] && s3[p] == s2[j-1]
	bp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				bp[i][j] = bp[i][j] || (bp[i-1][j] && s3[p] == s1[i-1])
			}
			if j > 0 {
				bp[i][j] = bp[i][j] || (bp[i][j-1] && s3[p] == s2[j-1])
			}
		}
	}
	return bp[n][m]
}

//leetcode submit region end(Prohibit modification and deletion)
