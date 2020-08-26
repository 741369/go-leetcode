/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午8:59
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
	two string
}

// ans 是答案
// one 代表第一个答案
type ans11 struct {
	one int
}

func Test_Problem11(t *testing.T) {

	qs := []question11{
		question11{
			para11{"abab", "ab"},
			ans11{0},
		},
		question11{
			para11{"aba", "ba"},
			ans11{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0459------------------------\n")

	for _, q := range qs {
		a, p := q.ans11, q.para11
		fmt.Printf("【input】:%v    【output】:%v      【true_output】:%v\n", p, violentMatch(p.one, p.two), a.one)
	}
}
