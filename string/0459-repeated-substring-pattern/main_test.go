/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/24 下午7:16
***********************************************/

package main

import (
	"fmt"
	"testing"
)

type question11 struct {
	para11
	ans11
}

// para 是参数
// one 代表第一个参数
type para11 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans11 struct {
	one bool
}

func Test_Problem11(t *testing.T) {

	qs := []question11{
		question11{
			para11{"abab"},
			ans11{true},
		},
		question11{
			para11{"aba"},
			ans11{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0459------------------------\n")

	for _, q := range qs {
		a, p := q.ans11, q.para11
		fmt.Printf("【input】:%v    【output】:%v      【true_output】:%v\n", p.one, repeatedSubstringPattern(p.one), a.one)
	}
}
