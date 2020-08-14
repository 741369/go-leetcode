/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/14 下午6:43
***********************************************/

package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	n := len(s)

	if n % 2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte
	for i := 0; i < n; i ++ {
		// 如果当前字符是右括号
		if pairs[s[i]] > 0 {
			// 如果栈内没有左括号，或者前一个括号不是当前括号的反括号，返回false
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 入栈左括号
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
//leetcode submit region end(Prohibit modification and deletion)

