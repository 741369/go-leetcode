/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/3/24 下午10:41
***********************************************/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {
		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}
