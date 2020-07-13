/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/13 上午9:35
***********************************************/

package main

import "fmt"

//给定两个数组，编写一个函数来计算它们的交集。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//
//
//
// 说明：
//
//
// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。
//
// Related Topics 排序 哈希表 双指针 二分查找
func main() {
	fmt.Printf("%#v\n", intersection([]int{1,2,2,1}, []int{2,2}))
	fmt.Printf("%#v\n", intersection([]int{4,9,5}, []int{9,4,9,8,4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) { // 减少空间复杂度，存储小的数组
		return intersection(nums2, nums1)
	}
	res := make([]int, 0)

	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}

	for _, v := range nums2 {
		if m[v] {
			res = append(res, v)
			m[v] = false
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
