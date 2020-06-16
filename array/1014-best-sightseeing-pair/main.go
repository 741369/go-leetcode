/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/17 上午1:22
***********************************************/

package main

import "fmt"

//给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
//
// 一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
//
// 返回一对观光景点能取得的最高分。
//
//
//
// 示例：
//
// 输入：[8,1,5,2,6]
//输出：11
//解释：i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11
//
//
//
//
// 提示：
//
//
// 2 <= A.length <= 50000
// 1 <= A[i] <= 1000
//
// Related Topics 数组

func main() {
	fmt.Println("===", maxScoreSightseeingPair([]int{8,1,5,2,6}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxScoreSightseeingPair(A []int) int {
	ans, mx := 0, A[0]+0
	for j := 1; j < len(A); j ++ {
		ans = max(ans, A[j] + mx - j)
		mx = max(mx, A[j]+j)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)

