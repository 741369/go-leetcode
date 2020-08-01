/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/1 下午11:26
***********************************************/

package main

import (
	"container/heap"
	"math"
)

//你有 k 个升序排列的整数数组。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
//
// 我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。
//
// 示例 1:
//
//
//输入:[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
//输出: [20,24]
//解释:
//列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
//列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
//列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
//
//
// 注意:
//
//
// 给定的列表可能包含重复元素，所以在这里升序表示 >= 。
// 1 <= k <= 3500
// -105 <= 元素的值 <= 105
// 对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。
//
// Related Topics 哈希表 双指针 字符串


//leetcode submit region begin(Prohibit modification and deletion)
var (
	next []int
	numsC [][]int
)

func smallestRange(nums [][]int) []int {
	numsC = nums
	rangeLeft, rangeRight := 0, math.MaxInt32
	minRange := rangeRight - rangeLeft
	max := math.MinInt32
	size := len(nums)
	next = make([]int, size)
	h := &IHeap{}
	heap.Init(h)

	for i := 0; i < size; i++ {
		heap.Push(h, i)
		max = Max(max, nums[i][0])
	}

	for {
		minIndex := heap.Pop(h).(int)
		curRange := max - nums[minIndex][next[minIndex]]
		if curRange < minRange {
			minRange = curRange
			rangeLeft, rangeRight = nums[minIndex][next[minIndex]], max
		}
		next[minIndex]++
		if next[minIndex] == len(nums[minIndex]) {
			break
		}
		heap.Push(h, minIndex)
		max = Max(max, nums[minIndex][next[minIndex]])
	}
	return []int{rangeLeft, rangeRight}
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return numsC[h[i]][next[h[i]]] < numsC[h[j]][next[h[j]]] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
