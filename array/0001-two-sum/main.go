/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/23 上午12:12
***********************************************/

package main

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)

	for k, v := range nums {
		if value, ok := tmp[target-v]; ok {
			return []int{value, k}
		}
		tmp[v] = k
	}
	return nil
}
