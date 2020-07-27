/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/27 下午8:58
***********************************************/

package main

import "fmt"

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"ae
//c"不是）。
//
// 示例 1:
//s = "abc", t = "ahbgdc"
//
// 返回 true.
//
// 示例 2:
//s = "axc", t = "ahbgdc"
//
// 返回 false.
//
// 后续挑战 :
//
// 如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码
//？
//
// 致谢:
//
// 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
// Related Topics 贪心算法 二分查找 动态规划
func main() {
	fmt.Println("======", isSubsequence("abc", "ahbgdc"))
	fmt.Println("======", isSubsequence("axc", "ahbgdc"))
	fmt.Println("======", isSubsequence("acb", "ahbgdc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0

	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i==n
}


func isSubsequence2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	cnt := 0
	j := 0
	for i := 0; i < len(s); i++ {
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				cnt++
				j++
				break
			}
		}
		if cnt == len(s) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
