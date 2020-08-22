/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/23 上午12:09
***********************************************/

package main

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start ++
		} else {
			high = height[end]
			end --
		}
		tmp := width * high
		if max < tmp {
			max = tmp
		}
	}
	return max
}

