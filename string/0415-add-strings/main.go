/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/30 上午11:01
***********************************************/

package main

import (
	"fmt"
)

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
// 注意：
//
//
// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
// Related Topics 字符串
func main() {
	fmt.Println(addStrings("193", "956"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	i, j, carry := len(num1)-1, len(num2)-1, 0
	res := ""
	for i >= 0 || j >= 0 {
		tmp := carry
		if i >= 0 {
			tmp += int(num1[i] - '0')
		}
		if j >= 0 {
			tmp += int(num2[j] - '0')
		}
		carry = tmp / 10
		res = fmt.Sprintf("%d%s", tmp%10, res)
		i--
		j--
	}
	if carry == 1 {
		res = fmt.Sprintf("%d%s", 1, res)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
