/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/10 下午8:33
***********************************************/

package main

import "fmt"

//给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
//
// 重复出现的子串要计算它们出现的次数。
//
// 示例 1 :
//
//
//输入: "00110011"
//输出: 6
//解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
//
//请注意，一些重复出现的子串要计算它们出现的次数。
//
//另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
//
//
// 示例 2 :
//
//
//输入: "10101"
//输出: 4
//解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
//
//
// 注意：
//
//
// s.length 在1到50,000之间。
// s 只包含“0”或“1”字符。
//
// Related Topics 字符串
func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countBinarySubstrings(s string) int {
	ptr, n, last, ans := 0, len(s), 0, 0
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		ans += min(last, count)
		last = count
	}
	return ans
}

func countBinarySubstrings2(s string) int {
	var counts []int
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr ++
			count ++
		}
		counts = append(counts, count)
	}

	ans := 0
	for i := 1; i < len(counts); i ++ {
		ans += min(counts[i], counts[i-1])
	}
	return ans
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
