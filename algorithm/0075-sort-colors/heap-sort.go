/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午10:14
***********************************************/
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
// 注意:
//不能使用代码库中的排序函数来解决这道题。
//
// 示例:
//
// 输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//
// 进阶：
//
//
// 一个直观的解决方案是使用计数排序的两趟扫描算法。
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
// Related Topics 排序 数组 双指针
package main

import "fmt"

func main() {
	nums := []int{0, 1}
	sortColors(nums)
	fmt.Println("======", nums)
}

func sortColors(nums []int) {
	heapSort(nums, len(nums))
}

func heapSort(nums []int, lens int) {
	buildHeap(nums, lens)
	for i := lens - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		lens--
		heap(nums, 0, lens)
	}
}

func buildHeap(nums []int, lens int) {
	for i := lens / 2; i >= 0; i-- {
		heap(nums, i, lens)
	}
}

func heap(nums []int, i int, lens int) {
	left, right, max := 2*i+1, 2*i+2, i
	if left < lens && nums[left] > nums[max] {
		max = left
	}
	if right < lens && nums[right] > nums[max] {
		max = right
	}
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heap(nums, max, lens)
	}
}
