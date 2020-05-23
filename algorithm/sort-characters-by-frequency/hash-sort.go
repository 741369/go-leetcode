/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午9:51
***********************************************/

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
