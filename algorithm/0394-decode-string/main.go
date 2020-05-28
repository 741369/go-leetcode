/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/28 下午12:54
***********************************************/

//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
// 示例:
//
//
//s = "3[a]2[bc]", 返回 "aaabcbc".
//s = "3[a2[c]]", 返回 "accaccacc".
//s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
//
// Related Topics 栈 深度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	str := "3[a]2[bc]"
	fmt.Println(decodeString(str))
}

func decodeString(s string) string {
	type pair struct {
		mul  int    //当前[]中字符串的倍数
		last string //是上个 [ 到当前 [ 的字符串
	}
	stack := list.New()
	var multi int = 0
	var res string = ""
	for _, v := range s {
		if v >= '0' && v <= '9' {
			multi = multi*10 + int(v-'0') //统计下一个[前的倍数
		} else if v == '[' {
			stack.PushBack(pair{multi, res}) //把当前[前所有的字符串组成的序列和当前[倍数压入栈
			multi, res = 0, ""
		} else if v == ']' {
			temp := stack.Back().Value.(pair)
			stack.Remove(stack.Back())
			res = temp.last + strings.Repeat(res, temp.mul) //字符串拼接
		} else {
			res += string(v) //遇到字母就直接拼接到当前res中直到遇到]时就进行
			//更新字符串序列
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
