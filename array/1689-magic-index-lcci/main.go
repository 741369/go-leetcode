/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/31 下午3:30
***********************************************/

package main

import "fmt"

//魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找
//出一个魔术索引，如果没有，则返回-1。若有多个魔术索引，返回索引值最小的一个。
//
// 示例1:
//
//  输入：nums = [0, 2, 3, 4, 5]
// 输出：0
// 说明: 0下标的元素为0
//
//
// 示例2:
//
//  输入：nums = [1, 1, 1]
// 输出：1
//
//
// 说明:
//
//
// nums长度在[1, 1000000]之间
// 此题为原书中的 Follow-up，即数组中可能包含重复元素的版本
//
// Related Topics 数组 二分查找
func main() {
	fmt.Println(findMagicIndex([]int{0, 2, 3, 4, 5}))
	fmt.Println(findMagicIndex([]int{1, 1, 1}))
}

func findMagicIndex(nums []int) int {
	return getAnswer(nums, 0, len(nums)-1)
}

func getAnswer(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (right + left) / 2
	leftAnswer := getAnswer(nums, left, mid-1)
	if leftAnswer != -1 {
		return leftAnswer
	} else if mid == nums[mid] {
		return mid
	}
	return getAnswer(nums, mid+1, right)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMagicIndex2(nums []int) int {
	for k, v := range nums {
		if k == v {
			return k
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
