/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/19 下午9:40
***********************************************/

package main

import "fmt"

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
//
//
// 示例 2:
//
// 输入: "race a car"
//输出: false
//
// Related Topics 双指针 字符串
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	alaNumStr := ""
	for i := 0; i < len(s); i++ {
		if isAlphabetOrNumber(s[i]) {
			if s[i] >= 'A' && s[i] <= 'Z' {
				alaNumStr += string(s[i] + 32)
			} else {
				alaNumStr += string(s[i])
			}
		}
	}

	lenNum := len(alaNumStr)
	for i := 0; i < lenNum/2; i++ {
		if alaNumStr[i] != alaNumStr[lenNum-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
func isAlphabetOrNumber(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
