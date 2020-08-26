/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午8:35
***********************************************/

package main

// Knuth-Morris-Pratt 字符串查找算法

// 问题：
// 有一个字符串S，和一个模式串P，现在要查找P在S中的为位置，怎么查找呢

// 暴力法
// 假设字符串S匹配到i位置，字符串P匹配到j位置，则有
// 1. 如果匹配成功，即S[i] = P[j] 则 i++, j++, 继续匹配下一个字符
// 2. 如果失败，即S[i] != P[j] ，则 i = i - j + 1, j = 0。即i回溯，j被置为0
func violentMatch(s string, p string) int {
	i, j, sLen, pLen := 0, 0, len(s), len(p)
	for i < sLen && j < pLen {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == pLen {
		return i-j
	}
	return -1
}

// https://blog.csdn.net/v_JULY_v/article/details/7041827
