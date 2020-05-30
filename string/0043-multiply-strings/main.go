/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/30 上午11:47
***********************************************/

package main

import "fmt"

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 示例 1:
//
// 输入: num1 = "2", num2 = "3"
//输出: "6"
//
// 示例 2:
//
// 输入: num1 = "123", num2 = "456"
//输出: "56088"
//
// 说明：
//
//
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
// Related Topics 数学 字符串
func main() {
	fmt.Println(multiply("123", "22"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			sum := n1*n2 + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	resStr := ""
	for k, v := range res {
		if k == 0 && v == 0 {
			continue
		}
		//resStr = fmt.Sprintf("%s%d", resStr, v)
		resStr += string(v + '0')
	}
	return resStr
}

func multiply2(num1 string, num2 string) string {
	res := "0"
	if num1 == "0" || num2 == "0" {
		return res
	}

	for i := len(num2) - 1; i >= 0; i-- {
		// 保存第i位num2和num乘积
		tmpMulti, carry := "", 0
		for j := 0; j < len(num2)-1-i; j++ {
			tmpMulti += "0"
		}
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0 || carry > 0; j-- {
			n1 := 0
			if j >= 0 {
				n1 = int(num1[j] - '0')
			}
			tmp := n1*n2 + carry
			tmpMulti = fmt.Sprintf("%d%s", tmp%10, tmpMulti)
			carry = tmp / 10
		}
		res = addStrings(res, tmpMulti)
	}
	return res
}

func addStrings(num1 string, num2 string) string {
	i, j, carry, res := len(num1)-1, len(num2)-1, 0, ""
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
