/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/23 下午9:21
***********************************************/

package main

import "fmt"

func main() {
	nums := []int{7, 3, 8, 5, 1, 2}
	//findKthLargest(nums, len(nums))
	fmt.Println("========", findKthLargest(nums, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	heapSort(nums, len(nums))
	return nums[k-1]
}

func heapSort(nums []int, lens int) {
	buildHeap(nums, lens)
	for i := lens - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		lens--
		heap(nums, 0, lens)
	}
}

// 构建最小堆
func buildHeap(nums []int, lens int) {
	for i := lens / 2; i >= 0; i-- {
		heap(nums, i, lens)
	}
}

func heap(nums []int, i int, lens int) {
	left, right, min := 2*i+1, 2*i+2, i
	if left < lens && nums[left] < nums[min] {
		min = left
	}
	if right < lens && nums[right] < nums[min] {
		min = right
	}
	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		heap(nums, min, lens)
	}
}
