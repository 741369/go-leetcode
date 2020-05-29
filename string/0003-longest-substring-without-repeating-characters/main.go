/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/29 下午9:31
***********************************************/
package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	tmp, ans := -1, 0
	for i := 0; i < len(s); i++ {
		// 指针i往前移动一次，map删除前一个元素
		if i > 0 {
			delete(m, s[i-1])
		}
		for tmp+1 < len(s) && m[s[tmp+1]] == 0 {
			m[s[tmp+1]]++
			tmp++
		}
		ans = max(tmp-i+1, ans)
	}
	return ans
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

//leetcode submit region end(Prohibit modification and deletion)

func lengthOfLongestSubstring2(s string) int {
	ans, left, right := 0, 0, 0
	tmpStr := s[left:right]
	for ; right < len(s); right++ {
		if index := strings.IndexByte(tmpStr, s[right]); index != -1 {
			left += index + 1
		}
		tmpStr = s[left : right+1]
		if len(tmpStr) > ans {
			ans = len(tmpStr)
		}
	}
	return ans
}
