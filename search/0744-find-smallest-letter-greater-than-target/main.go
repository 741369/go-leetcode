/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/24 下午5:15
***********************************************/
//给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。
//
// 在比较时，字母是依序循环出现的。举个例子：
// 如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'
// 示例：
//
// 输入:
//letters = ["c", "f", "j"]
//target = "a"
//输出: "c"
//
//输入:
//letters = ["c", "f", "j"]
//target = "c"
//输出: "f"
//
//输入:
//letters = ["c", "f", "j"]
//target = "d"
//输出: "f"
//
//输入:
//letters = ["c", "f", "j"]
//target = "g"
//输出: "j"
//
//输入:
//letters = ["c", "f", "j"]
//target = "j"
//输出: "c"
//
//输入:
//letters = ["c", "f", "j"]
//target = "k"
//输出: "c"
//
//
//
//
// 提示：
//
//
// letters长度范围在[2, 10000]区间内。
// letters 仅由小写字母组成，最少包含两个不同的字母。
// 目标字母target 是一个小写字母。
//
// Related Topics 二分查找
package main

import "fmt"

func main() {
	fmt.Printf("%v", string(nextGreatestLetter([]byte{'a', 'b', 'c'}, 'd')))
}

// 比target大的最小字母，如果都不比target大的话就返回第一个。
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		middle := left + (right-left)/2
		if letters[middle] <= target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return letters[left%(len(letters))]
}
