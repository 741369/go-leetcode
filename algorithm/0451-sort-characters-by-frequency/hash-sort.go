/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午9:51
***********************************************/
//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
// 示例 1:
//
//
//输入:
//"tree"
//
//输出:
//"eert"
//
//解释:
//'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//
//
// 示例 2:
//
//
//输入:
//"cccaaa"
//
//输出:
//"cccaaa"
//
//解释:
//'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//
//
// 示例 3:
//
//
//输入:
//"Aabb"
//
//输出:
//"bbAa"
//
//解释:
//此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
//
// Related Topics 堆 哈希表
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "aaaaccabb"
	fmt.Println("==", frequencySort(str))
}

func frequencySort(s string) string {
	// 用map收集相同字符
	strs := strings.Split(s, "")
	strsMap := make(map[string]string)
	for i := 0; i < len(strs); i++ {
		strsMap[strs[i]] += strs[i]
	}
	// 接下来要做的就是堆strsMap的value进行排序，不一定要用二叉堆来维护最值
	strList := make([]string, 0, len(strsMap))
	for _, val := range strsMap {
		strList = append(strList, val)
	}
	sort.Slice(strList, func(i, j int) bool {
		return len(strList[i]) >= len(strList[j])
	})
	fmt.Printf("======%#v\n", strsMap)
	fmt.Printf("======%#v\n", strList)
	// 返回拼接的字符串
	return strings.Join(strList, "")
}
