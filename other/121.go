/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/3/9 下午9:34
***********************************************/

package main

import "fmt"

//将股票价格数组转化为股票涨跌数组，则问题转化为求最大子数组和问题
//动态规划解法，时间复杂度O(n),空间复杂度O(n)
/**
输入: [7,1,5,3,6,4]
输出: 5
*/
func maxProfit2(prices []int) int {

	l := len(prices)
	if l == 0 {
		return 0
	}
	max := 0
	nums := make([]int, l)

	nums[0] = 0
	for i := 1; i < l; i++ {
		nums[i] = prices[i] - prices[i-1]
	}

	fmt.Println("=====", nums)

	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		fmt.Println("===", i, "==", nums)
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

/**
输入: [7,1,5,3,6,4]
输出: 5
i = 1 max = 0, min = 1
i = 2 max = 4, min = 1
i = 3 max = 4, min = 1
i = 4 max = 5, min = 1
i = 5 max = 5, min = 1
*/
func maxProfit(prices []int) int { // O(n)
	if len(prices) == 0 {
		return 0
	}
	Min := prices[0] // 记录当前股票最小价格
	Max := 0         // 记录当前最大获利
	for i := 1; i < len(prices); i++ {
		if prices[i]-Min > Max {
			Max = prices[i] - Min
		}
		if prices[i] < Min {
			Min = prices[i]
		}
		fmt.Println("=====", Max, "==", Min)
	}
	return Max
}

func main() {
	max := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println("==", max)
}
