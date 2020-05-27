/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/15 下午8:11
***********************************************/

package main

import (
	"fmt"
)

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{-2, 0, 1, 1, 2}
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum2(nums []int) [][]int {
	sumArr := make([][]int, 0)

	mapTmp := make(map[[3]int]bool, len(nums))
	sortArr(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-2; i++ {
		twoSumMap := make(map[int]bool, len(nums))
		for j := i + 1; j < len(nums); j++ {
			tmp := 0 - nums[i] - nums[j]
			if twoSumMap[tmp] {
				if !mapTmp[[3]int{nums[i], tmp, nums[j]}] {
					sumArr = append(sumArr, []int{nums[i], tmp, nums[j]})
				}
				mapTmp[[3]int{nums[i], tmp, nums[j]}] = true
			} else {
				twoSumMap[nums[j]] = true
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

/**
 * @Author liu
 * @Description // 方法二
 * @Date 2020/5/15 下午11:22
 * @Param
 * @return
 **/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sumArr := make([][]int, 0)
	sortArr(nums, 0, len(nums)-1)
	if nums[0] <= 0 && nums[len(nums)-1] >= 0 {
		for k := 0; k < len(nums)-2; k++ {
			// 如果最左元素大于0，肯定没有满足条件的
			if nums[k] > 0 {
				break
			}
			// 去重
			if k > 0 && nums[k] == nums[k-1] {
				continue
			}
			left, right := k+1, len(nums)-1
			for left < right {
				target := nums[k] + nums[left] + nums[right]
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
	}
	return sumArr
}
