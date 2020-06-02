/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/31 下午6:23
***********************************************/

package main

import "fmt"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
func main() {
	fmt.Printf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortArr(nums []int, left int, right int) {
	if left < right {
		p := partition(nums, left, right)
		sortArr(nums, left, p-1)
		sortArr(nums, p+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	p, i, j := nums[left], left, right
	for i < j {
		for i < j && p < nums[j] {
			j--
		}
		for i < j && nums[i] <= p {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

// 双指针前后移动
func threeSum(nums []int) [][]int {
	sumArr := make([][]int, 0)

	// 1. 排序，判断边界条件
	sortArr(nums, 0, len(nums)-1)
	if len(nums) < 3 {
		return nil
	}

	if nums[0] <= 0 && nums[len(nums)-1] >= 0 {
		// 2. 从第一个元素到倒数第二个元素遍历，加上left，right两个指针构成3个元素
		for k := 0; k < len(nums)-2; k++ {
			// 判断边界条件，第一个元素大于0则不满足条件
			if nums[k] > 0 {
				break
			}
			// 去重,和前面的元素值相同的可以过滤掉
			if k > 0 && nums[k] == nums[k-1] {
				continue
			}
			left, right := k+1, len(nums)-1
			for left < right {
				// 3. 左右指针元素和遍历元素之和满足条件则添加成功队列，不满足则指针左右移动
				target := nums[k] + nums[left] + nums[right]
				if target == 0 {
					sumArr = append(sumArr, []int{nums[k], nums[left], nums[right]})
					// 对满足条件的元素，相同的去重
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if target < 0 {
					left++
				} else {
					right--
				}
			}

		}
	}
	return sumArr
}

//leetcode submit region end(Prohibit modification and deletion)
