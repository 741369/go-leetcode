/**********************************************
** @Des: 820 short-encoding-of-words
** @Author: 1@lg1024.com
** @Last Modified time: 2020/3/28 下午12:14
***********************************************/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "a"
	c := "你"
	r := []rune(c)
	//fmt.Println(len(c), len(s)) // 1 3
	fmt.Printf("%x, %v \n", s, s)
	fmt.Printf("%x, %v \n", c, c)

	fmt.Println("==", r)

	str := "你你"

	for _, s := range str {
		fmt.Printf("unicode: %c %d\n", s, s)
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("ascii: %c %d\n", str[i], str[i])
	}

	// ASCII 字符串长度使用len() 长度
	//Unicode 字符串长度使用utf8.RuneCountInString()
	fmt.Println("===", len(str))
	fmt.Println("===", len([]rune(str)))
	fmt.Println("===", utf8.RuneCountInString(str))

	//words := []string{"time", "me", "bell"}
	//fmt.Println(minimumLengthEncoding(words))
}

func minimumLengthEncoding(words []string) int {
	wordsMap := make(map[string]string, len(words))
	for _, v := range words {
		wordsMap[v] = v
	}
	for _, v := range words {
		for i := 1; i < len(v); i++ {
			if wordsMap[v[i:]] != "" {
				wordsMap[v[i:]] = ""
			}
		}
	}

	ans := 0
	for _, v := range wordsMap {
		if v != "" {
			ans += len(v) + 1
		}
	}
	return ans
}
