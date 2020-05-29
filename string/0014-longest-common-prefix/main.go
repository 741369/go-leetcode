/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/29 下午10:03
***********************************************/

package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	ans, minLength := "", len(str[0])
	for i := 1; i < len(str); i++ {
		if len(str[i]) < minLength {
			minLength = len(str[i])
		}
	}

	for i := 0; i < minLength; i++ {
		tmp := str[0][i]
		for j := 1; j < len(str); j++ {
			if tmp != str[j][i] {
				return ans
			}
		}
		ans += string(tmp)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
