/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午8:11
***********************************************/

package main

import "fmt"

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 示例 1:
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
// 说明:
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Printf("%d\n", findKthLargest(nums, 2))
	p := quickSort(nums, 2, 0, len(nums)-1)
	fmt.Printf("%#v, %d\n", nums, p)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	return quickSort(nums, k-1, 0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

// 快速排序方法
func quickSort(nums []int, k int, left int, right int) int {
	if left < right {
		p := partition(nums, left, right)
		if p == k {
			return nums[p]
		}
		quickSort(nums, k, left, p-1)
		quickSort(nums, k, p+1, right)
	}
	return nums[k]
}
func partition(nums []int, left int, right int) int {
	i, j, temp := left, right, nums[left]
	for i < j {
		for i < j && temp > nums[j] {
			j--
		}
		for i < j && nums[i] >= temp {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

// 堆排序方法
