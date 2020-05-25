/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/25 上午9:01
***********************************************/
//给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 *
// 。
//
// 示例 1:
//
// 输入: "2-1-1"
//输出: [0, 2]
//解释:
//((2-1)-1) = 0
//(2-(1-1)) = 2
//
// 示例 2:
//
// 输入: "2*3-4*5"
//输出: [-34, -14, -10, -10, 10]
//解释:
//(2*(3-(4*5))) = -34
//((2*3)-(4*5)) = -14
//((2*(3-4))*5) = -10
//(2*((3-4)*5)) = -10
//(((2*3)-4)*5) = 10
// Related Topics 分治算法

package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "2*3-4*5"
	fmt.Println("=====", diffWaysToCompute(str))
}

func diffWaysToCompute(str string) []int {
	var res []int

	for i, v := range []byte(str) {
		if v != '+' && v != '-' && v != '*' {
			continue
		}
		leftStr := diffWaysToCompute(str[:i])    // 不包括i
		rightStr := diffWaysToCompute(str[i+1:]) // 包括i+1
		for _, leftV := range leftStr {
			for _, rightV := range rightStr {
				switch v {
				case '+':
					res = append(res, leftV+rightV)
				case '-':
					res = append(res, leftV-rightV)
				case '*':
					res = append(res, leftV*rightV)
				}
			}
		}
	}
	if len(res) == 0 {
		v, _ := strconv.Atoi(str)
		res = append(res, v)
	}
	return res
}
