/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/8 下午11:55
***********************************************/

package main

import "fmt"

//你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方
//法，生成跳水板所有可能的长度。
// 返回的长度需要从小到大排列。
// 示例：
// 输入：
//shorter = 1
//longer = 2
//k = 3
//输出： {3,4,5,6}
//
// 提示：
//
// 0 < shorter <= longer
// 0 <= k <= 100000
//
// Related Topics 递归 记忆化
func main() {
	fmt.Printf("%#v", divingBoard(1, 2, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{k * shorter}
	}
	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = shorter*(k-i) + longer*i
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
