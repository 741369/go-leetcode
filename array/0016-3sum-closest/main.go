/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/24 下午7:07
***********************************************/

package main

import (
	"fmt"
	"math"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
func main() {
	fmt.Println("========", threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println("========", threeSumClosest([]int{1, 1, 1, 0}, -100))
}

//leetcode submit region begin(Prohibit modification and deletion)
//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	lenNums := len(nums) - 1
	quickSort(nums, 0, len(nums)-1)
	tmp, sum := math.MaxInt64, 0

	for k := 0; k < lenNums-1; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i, j := k+1, lenNums
		for i < j {
			min, flag := minDiff(target, nums[k]+nums[i]+nums[j])
			if min == 0 {
				return nums[k] + nums[i] + nums[j]
			} else if min < tmp {
				tmp = min
				sum = nums[k] + nums[i] + nums[j]
			}
			if flag {
				i++
			} else {
				j--
			}
		}
	}
	return sum
}

// 绝对值
func minDiff(i, j int) (int, bool) {
	if i > j {
		return i - j, true
	}
	return j - i, false
}

//leetcode submit region end(Prohibit modification and deletion)

func quickSort(nums []int, left int, right int) {
	if left < right {
		p := partition(nums, left, right)
		quickSort(nums, left, p-1)
		quickSort(nums, p+1, right)
	}
}

// 找到指定元素i，i的左边都比i的值小，右边都比i的值大
func partition(nums []int, left int, right int) int {
	pas, i, j := nums[left], left, right
	for i < j {
		for i < j && pas < nums[j] {
			j--
		}
		for i < j && nums[i] <= pas { // 加上等于可以减少交换次数
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i

}
