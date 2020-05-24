/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/24 下午5:07
***********************************************/
//实现 int sqrt(int x) 函数。
//
// 计算并返回 x 的平方根，其中 x 是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
// 示例 1:
//
// 输入: 4
//输出: 2
//
//
// 示例 2:
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
// Related Topics 数学 二分查找
package main

import "fmt"

func main() {
	fmt.Println(mySqrt(15))
}

func mySqrt(x int) int {
	left, right, ans := 0, x, -1
	for left <= right {
		middle := left + (right-left)/2
		if middle*middle <= x {
			ans = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return ans
}
