/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/24 下午7:16
***********************************************/

package main

import (
	"fmt"
	"strings"
)

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。
//给定的字符串只含有小写英文字母，并且长度不超过10000。
//
// 示例 1:
//
//
//输入: "abab"
//
//输出: True
//
//解释: 可由子字符串 "ab" 重复两次构成。
//
//
// 示例 2:
//
//
//输入: "aba"
//
//输出: False
//
//
// 示例 3:
//
//
//输入: "abcabcabcabc"
//
//输出: True
//
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
//
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)

func repeatedSubstringPattern(s string) bool {
	return kmp(s)
}

func kmp(pattern string) bool {
	n := len(pattern)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i - 1]
		for j != -1 && pattern[j + 1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j + 1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	return fail[n - 1] != -1 && n % (n - fail[n - 1] - 1) == 0
}

func repeatedSubstringPattern3(s string) bool {
	fmt.Println("0"+(s+s)[1:], "=========", strings.Index((s + s)[1:], s))
	return strings.Index("0"+(s+s)[1:], s) != len(s)
}

func repeatedSubstringPattern2(s string) bool {
	n := len(s)
	for i := 1; i * 2 <= n; i++ {
		if n % i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j - i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
