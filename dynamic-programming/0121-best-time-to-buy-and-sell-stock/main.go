/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/7 下午7:50
***********************************************/

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
// 注意：你不能在买入股票前卖出股票。
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
//
// 示例 2:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
// Related Topics 数组 动态规划
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("========", maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println("========", maxProfit([]int{7,6,4,3,1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]  // 记录当前股票最小价格
	max := 0  // 记录最大获利
	for i := 1; i < len(prices); i ++ {
		if prices[i] < min {
			min = prices[i]
		}
		if max < (prices[i] - min) {
			max = prices[i] - min
		}
	}
	return max
}
//leetcode submit region end(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	dpi0 := 0
	dpi1 := math.MinInt64
	for i := 0; i < len(prices); i ++ {
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, -prices[i])
	}
	return dpi0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}