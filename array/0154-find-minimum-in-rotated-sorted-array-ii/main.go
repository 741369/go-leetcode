package main
// 154. 寻找旋转排序数组中的最小值 II
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 请找出其中最小的元素。

// 注意数组中可能存在重复的元素。

// 示例 1：

// 输入: [1,3,5]
// 输出: 1
// 示例 2：

// 输入: [2,2,2,0,1]
// 输出: 0
// 说明：

// 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
// 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？

import "fmt"

func main() {
	fmt.Println(findMin([]int{1,3,5}))
	fmt.Println(findMin([]int{2,2,2,0,1}))
}
// 二分查找
func findMin(nums []int) int {
	left, right := 0, len(nums) - 1
	// 将mid和right比较有三种情况
	// 1. mid < high 说明最小值在最左边
	// 2. mid > high 说明最小值在最右边
	// 3. mid = high 说明不确定在哪边，可以hght --
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right --
		}
	}
	return nums[left]
}

// 直接遍历
func findMin(nums []int) int {
	for i := 0; i < len(nums)-1; i++{
		if nums[i+1] < nums[i] {
			return nums[i+1]
		}
	}
	return nums[0]
}