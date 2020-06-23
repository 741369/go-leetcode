/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/23 上午12:43
***********************************************/

package main

import (
	"fmt"
	"strconv"
)

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
// 输入: a = "11", b = "1"
//输出: "100"
//
// 示例 2:
//
// 输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//
// 提示：
//
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
//
// Related Topics 数学 字符串
func main() {
	fmt.Println("=====", addBinary("1010", "1011"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	aL, bL, plus, res := len(a)-1, len(b)-1, 0, ""
	for plus > 0 || (aL >= 0 && bL >= 0) {
		tmp := plus
		if aL >= 0 {
			tmp += int(a[aL] - '0')
		}
		if bL >= 0 {
			tmp += int(b[bL] - '0')
		}
		res = strconv.Itoa(tmp % 2) + res
		plus = tmp / 2

		//fmt.Println("=======", a[aL] - '0', "==", b[bL] - '0', "===", tmp % 2, "==", tmp / 2, "==", strconv.Itoa(tmp % 2), "==", res)

		aL --
		bL --
	}
	if aL >= 0 {
		res = a[:aL+1] + res
	}
	if bL >= 0 {
		res = b[:bL+1] + res
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
