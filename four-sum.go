/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/16 上午12:08
***********************************************/

package main

import "fmt"

func main() {
	//nums := []int{0, 0, 0, 0}  0
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{5, 5, 3, 5, 1, -5, 1, -2}
	//nums := []int{1, 0, -1, 0, -2, 2}  0
	fmt.Println(fourSum(nums, 4))
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	// 排序
	sortArr(nums, 0, len(nums)-1)
	sumArr := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum := threeSum(nums[i+1:], target-nums[i])
		for _, v := range threeSum {
			sumArr = append(sumArr, append(v, nums[i]))
		}
	}
	return sumArr
}

func threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sumArr := make([][]int, 0)
	for k := 0; k < len(nums)-2; k++ {
		// 去重
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		left, right := k+1, len(nums)-1
		for left < right {
			target := nums[k] + nums[left] + nums[right] - target
			if target == 0 {
				sumArr = append(sumArr, []int{nums[k], nums[left], nums[right]})
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
	return sumArr
}

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
		for i < j && p <= nums[j] {
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
