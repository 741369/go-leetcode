/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/11 下午11:02
***********************************************/

package main

import "fmt"

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//
// 示例 1:
//
// 输入: 2.00000, 10
//输出: 1024.00000
//
//
// 示例 2:
//
// 输入: 2.10000, 3
//输出: 9.26100
//
//
// 示例 3:
//
// 输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25
//
// 说明:
//
//
// -100.0 < x < 100.0
// n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
//
// Related Topics 数学 二分查找

func main() {
	fmt.Println(myPow(2.0000, 10))
	fmt.Println(myPow(2.1000, 3))
	fmt.Println(myPow(2.0000, -2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	// 递归终止条件
	if n == 0 {
		return 1
	}
	// 递归
	y := quickMul(x, n / 2)
	// 逻辑处理
	if n % 2 == 0 {
		return  y * y
	}
	return y * y * x
}

//leetcode submit region end(Prohibit modification and deletion)

