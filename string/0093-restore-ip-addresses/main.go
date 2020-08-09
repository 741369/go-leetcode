/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/30 下午9:10
***********************************************/

package main

import "fmt"

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
// 有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。
// 示例:
//
// 输入: "25525511135"
//输出: ["255.255.11.135", "255.255.111.35"]
// Related Topics 字符串 回溯算法
func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	// 判断边界条件
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	var result []string
	// 1. 第一个小数点后，最多3位，并且保证后面最少有3位，
	for i := 1; i < 4 && i < len(s)-2; i++ {
		for j := i + 1; j < i+4 && j < len(s)-1; j++ {
			for k := j + 1; k < j+4 && k < len(s); k++ {
				if judgeNumber(s[:i]) && judgeNumber(s[i:j]) && judgeNumber(s[j:k]) && judgeNumber(s[k:]) {
					result = append(result, s[:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:])
				}
			}
		}
	}
	return result
}

// 判断数字是否为ip地址数字1-255
func judgeNumber(s string) bool {

	// 第一位不能是0，也排除了数字0
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	// 数字大小范围为1, 255
	result := 0
	for i := 0; i < len(s); i++ {
		result = result*10 + int(s[i]-'0')
	}
	if result > 255 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
