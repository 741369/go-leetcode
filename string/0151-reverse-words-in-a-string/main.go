/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/30 下午5:04
***********************************************/

package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，逐个翻转字符串中的每个单词。
//
// 示例 1：
//
// 输入: "the sky is blue"
//输出: "blue is sky the"
//
// 示例 2：
//
// 输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
// 示例 3：
//
// 输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
// 说明：
//
// 无空格字符构成一个单词。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
// 进阶：
//
// 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。
// Related Topics 字符串
func main() {
	fmt.Println(reverseWords("a b c     ddddd"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	queue, word := *new([]string), *new([]byte)

	for k, v := range s {
		if v == ' ' {
			if len(word) > 0 {
				queue = append(queue, string(word))
				word = []byte{}
			}
		} else {
			word = append(word, s[k])
		}
	}
	if len(word) > 0 {
		queue = append(queue, string(word))
	}
	res := ""
	for k, v := range queue {
		if k == 0 {
			res = v
		} else {
			res = fmt.Sprintf("%s %s", v, res)
		}
	}
	return res
}

func reverseWords3(s string) string {
	// 1. 通过空格将字符切分为数组
	// 2. 将数组反转
	// 3. 通过空格组合新数组为字符串
	arr := strings.Split(s, " ")
	var newArr []string
	for i := len(arr) - 1; i >= 0; i-- {
		if strings.TrimSpace(arr[i]) != "" {
			newArr = append(newArr, arr[i])
		}
	}
	return strings.Join(newArr, " ")
}

func reverseWords2(s string) string {
	list := strings.Split(s, " ")
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		if len(list[i]) > 0 {
			res = append(res, list[i])
		}
	}
	s = strings.Join(res, " ")
	return s
}

//leetcode submit region end(Prohibit modification and deletion)
