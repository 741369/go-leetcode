/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/30 上午12:05
***********************************************/

package main

import (
	"fmt"
	"strings"
)

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
// 换句话说，第一个字符串的排列之一是第二个字符串的子串。
// 示例1:
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
// 示例2:
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
// 注意：
// 输入的字符串只包含小写字母
// 两个字符串的长度都在 [1, 10,000] 之间
//
// Related Topics 双指针 Sliding Window
func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	res := permutation(s1)
	for k := range res {
		if strings.Contains(s2, k) {
			return true
		}
	}
	return false
}

func permutation(s string) map[string]struct{} {
	m := map[string]struct{}{}
	dfs(0, []byte(s), m)
	return m
}

// 深度优先
func dfs(start int, s []byte, m map[string]struct{}) {
	if start == len(s)-1 {
		m[string(s)] = struct{}{} // 保存符合条件的字符
		return
	}
	tmp := map[byte]bool{}
	for i := start; i < len(s); i++ {
		if tmp[s[i]] { // 剪枝
			continue
		}

		tmp[s[i]] = true
		s[i], s[start] = s[start], s[i] // 交换，将 s[i] 固定在第 start 位
		dfs(start+1, s, m)              // 开启固定第start+1位字符
		s[i], s[start] = s[start], s[i] // 恢复交换
	}
}

//方法1
func checkInclusion2(s1 string, s2 string) bool {
	res := permutation2(s1)
	for k := range res {
		if strings.Contains(s2, k) {
			return true
		}
	}
	return false
}

func permutation2(s string) map[string]bool {
	res := map[string]bool{}
	bfs(s, "", res)
	return res
}

func bfs(remained string, now string, m map[string]bool) {
	if len(remained) == 0 {
		m[now] = true
	}
	for i := 0; i < len(remained); i++ {
		tmp := remained[:i] + remained[i+1:]
		bfs(tmp, now+string(remained[i]), m)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
