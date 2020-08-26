/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/25 下午11:36
***********************************************/

package main

import "math"

//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//
// 示例:
//
//
//输入: [4, 6, 7, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7
//]]
//
// 说明:
//
//
// 给定数组的长度不会超过15。
// 数组中的整数范围是 [-100,100]。
// 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
//
// Related Topics 深度优先搜索


//leetcode submit region begin(Prohibit modification and deletion)
var (
	n int
	temp []int
)
// 枚举 + 哈希
func findSubsequences(nums []int) [][]int {
	n = len(nums)
	var ans [][]int
	set := map[int]bool{}
	for i := 0; i < 1 << n; i++ {
		findSubsequences1(i, nums)
		hashValue := getHash(263, int(1e9 + 7))
		if check() && !set[hashValue] {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
			set[hashValue] = true
		}
	}
	return ans
}

func findSubsequences1(mask int, nums []int) {
	temp = []int{}
	for i := 0; i < n; i++ {
		if (mask & 1) != 0 {
			temp = append(temp, nums[i])
		}
		mask >>= 1
	}
}

func getHash(base, mod int) int {
	hashValue := 0
	for _, x := range temp {
		hashValue = hashValue * base % mod + (x + 101)
		hashValue %= mod
	}
	return hashValue
}

func check() bool {
	for i := 1; i < len(temp); i++ {
		if temp[i] < temp[i - 1] {
			return false
		}
	}
	return len(temp) >= 2
}

//leetcode submit region end(Prohibit modification and deletion)

var (
	temp2 []int
	ans [][]int
)

func findSubsequences2(nums []int) [][]int {
	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

func dfs(cur, last int, nums []int) {
	if cur == len(nums) {
		if len(temp2) >= 2 {
			t := make([]int, len(temp2))
			copy(t, temp2)
			ans = append(ans, t)
		}
		return
	}
	if nums[cur] >= last {
		temp2 = append(temp2, nums[cur])
		dfs(cur + 1, nums[cur], nums)
		temp2 = temp[:len(temp2)-1]
	}
	if nums[cur] != last {
		dfs(cur + 1, last, nums)
	}
}

