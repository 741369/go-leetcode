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
	fmt.Println(checkInclusion("adc", "dcda"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 优化的滑动窗口
func checkInclusion7(s1 string, s2 string) bool {
	cnt1, cnt2 := [26]int{}, [26]int{}

	for i := range s1 {
		cnt1[s1[i]-'a']++
	}

	start := 0
	for j := range s2 {
		cnt2[s2[j]-'a']++
		for start <= j && cnt2[s2[start]-'a'] > cnt1[s2[start]-'a'] {
			cnt2[s2[start]-'a']--
			start++
		}
		if cnt1 == cnt2 {
			return true
		}
	}
	return cnt1 == cnt2
}

// 更优化的滑动窗口，减少比较次数，cnt1 和 cnt2 一起初始化
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++
	}
	return cnt1 == cnt2
}

// 简单的滑动窗口
func checkInclusion5(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m1, m2 := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
		m2[s2[i]]++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if match(m1, m2) {
			return true
		}
		fmt.Println("==", m1, "==", m2)
		m2[s2[i]]--
		m2[s2[i+len(s1)]]++
		fmt.Println("==", "==", m2)
	}
	return match(m1, m2)
}

func match(m1, m2 map[byte]int) bool {
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func checkInclusion4(s1 string, s2 string) bool {
	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Tmp := s2[i : i+len(s1)]
		for _, v := range []byte(s1) {
			if index := strings.IndexByte(s2Tmp, v); index == -1 {
				fmt.Println("=====", s2Tmp, "==", v, "==", index)
				break
			} else {
				tmp := []byte(s2Tmp)
				s2Tmp = string(tmp[:index]) + string(tmp[index+1:])
			}
		}
		fmt.Println("===========", s2Tmp)
		if s2Tmp == "" {
			return true
		}
	}
	return false
}

func checkInclusion3(s1 string, s2 string) bool {
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
