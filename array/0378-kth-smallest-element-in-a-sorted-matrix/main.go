/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/2 下午11:41
***********************************************/

package main

import (
	"fmt"
	"sort"
)

//给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
//
//
//
// 示例：
//
// matrix = [
//   [ 1,  5,  9],
//   [10, 11, 13],
//   [12, 13, 15]
//],
//k = 8,
//
//返回 13。
//
//
//
//
// 提示：
//你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。
// Related Topics 堆 二分查找
func main() {
	matrix := [][]int{
		{1,  5,  9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(matrix, 8))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 直接排序,
// O(n^2logn), 对n^2个数组排序
// O(n^2)
func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	rows, columns := len(matrix), len(matrix[0])
	sorted := make([]int, rows * columns)
	index := 0
	for _, row := range matrix {
		for _, column := range row {
			sorted[index] = column
			index ++
		}
	}
	sort.Ints(sorted)
	return sorted[k-1]
}
//leetcode submit region end(Prohibit modification and deletion)

