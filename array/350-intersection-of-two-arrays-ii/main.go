/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/13 上午9:05
***********************************************/

package main

import (
	"fmt"
	"sort"
)

//给定两个数组，编写一个函数来计算它们的交集。
//
// 示例 1:
//
// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2,2]
//
//
// 示例 2:
//
// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9]
//
// 说明：
//
//
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
// 我们可以不考虑输出结果的顺序。
//
//
// 进阶:
//
//
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
//
// Related Topics 排序 哈希表 双指针 二分查找
func main() {
	fmt.Printf("%#v\n", intersect([]int{1,2,2,1}, []int{2,2}))
	fmt.Printf("%#v\n", intersect([]int{4,9,5}, []int{9,4,9,8,4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 排序双指针法
func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i ++
			j ++
		} else if nums1[i] < nums2[j] {
			i ++
		} else {
			j ++
		}
	}

	return res
}


// 哈希法
func intersect2(nums1 []int, nums2 []int) []int {
	// 为了降低空间复杂度，首先遍历较短的数组并在哈希表中记录每个数字以及对应出现的次数，然后遍历较长的数组得到交集。
	if len(nums1) > len(nums2) {
		intersect2(nums2, nums1)
	}
	res := make([]int, 0)
	m := make(map[int]int)

	for _, v := range nums1 {
		m[v] ++
	}

	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v] --
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
