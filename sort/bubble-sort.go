/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/11 下午9:43
***********************************************/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	arr := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("====", BubbleSort(arr))
	fmt.Println("====", InsertSort(arr))
	fmt.Println("====", SelectionSort(arr))
	fmt.Println("====", ShellSort(arr))
	fmt.Println("====", QuickSort(arr))
	fmt.Println("====", HeadSort(arr))
	fmt.Println("====", MergeSort(arr))
	fmt.Println("====", CountSort(arr))
	fmt.Println("====", BinSort(arr))
	fmt.Println("====", RadixSort(arr))
	fmt.Println("====", TopK(arr, 5))
}

// 冒泡排序
func BubbleSort(arr []int) []int {
	if len(arr) < 1 {
		return nil
	}

	for i := 0; i < len(arr); i++ {
		flag := false

		// 每次循环，将最大的元素移动到最后
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// 插入排序
func InsertSort(arr []int) []int {
	if len(arr) < 1 {
		return nil
	}

	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1

		// 每次循环查询temp要插入的位置
		for j >= 0 && temp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}

// 选择排序
func SelectionSort(arr []int) []int {
	if len(arr) < 1 {
		return nil
	}

	for i := 0; i < len(arr)-1; i++ {
		// 查找最小的索引
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换最小的元素到最前面
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 希尔排序
func ShellSort(arr []int) []int {
	step := len(arr) / 2
	for ; step > 0; step = step / 2 {
		for i := step; i < len(arr); i++ {
			temp := arr[i]
			j := i - step
			// 每次循环查询temp要插入的位置
			for j >= 0 && temp < arr[j] {
				arr[j+step] = arr[j]
				j = j - step
			}
			arr[j+step] = temp
		}
	}
	return arr
}

// 快速排序
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, left, right int) {
	if left < right {
		index := arr[left]
		low, high := left, right
		for low < high {
			for low < high && index <= arr[high] {
				high--
			}
			arr[low] = arr[high]
			for low < high && index >= arr[low] {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = index
		quickSort(arr, left, low)
		quickSort(arr, low+1, right)
	}
}

// 堆排序
func HeadSort(arr []int) []int {
	for i := len(arr) / 2; i >= 0; i-- {
		initHead(arr, i, len(arr))
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		initHead(arr, 0, i)
	}
	return arr
}
func initHead(nums []int, parent, len int) {
	temp := nums[parent]
	child := 2*parent + 1

	for child < len {
		if child+1 < len && nums[child] < nums[child+1] {
			child++
		}
		if child < len && nums[child] <= temp {
			break
		}
		nums[parent] = nums[child]
		parent = child
		child = child*2 + 1
	}
	nums[parent] = temp
}

// 归并排序
func MergeSort(arr []int) []int {
	mergeSort(arr, 0, len(arr)-1)
	return arr
}
func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}
func merge(arr []int, left, mid, right int) {
	i := left
	j := mid + 1
	var tmp []int
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	if i <= mid {
		tmp = append(tmp, arr[i:mid+1]...)
	} else {
		tmp = append(tmp, arr[j:right+1]...)
	}
	for k := 0; k < len(tmp); k++ {
		arr[left+k] = tmp[k]
	}
}

// 计数排序
func CountSort(arr []int) []int {
	// 求数组的最大值
	maxNum := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxNum < arr[i] {
			maxNum = arr[i]
		}
	}

	arrTemp := make([]int, maxNum+1)
	for j := 0; j < len(arr); j++ {
		arrTemp[arr[j]]++
	}

	k := 0
	for m, n := range arrTemp {
		for p := 0; p < n; p++ {
			arr[k] = m
			k++
		}
	}
	return arr
}

// 桶排序
func BinSort(arr []int) []int {
	minNum, maxNum, binNum := arr[0], arr[0], len(arr)
	// 求数组最大值，最小值
	for i := 0; i < len(arr); i++ {
		if minNum > arr[i] {
			minNum = arr[i]
		}
		if maxNum < arr[i] {
			maxNum = arr[i]
		}
	}

	// 将对数放入对应的桶
	bin := make([][]int, binNum)
	for j := 0; j < len(arr); j++ {
		n := (arr[j] - minNum) / ((maxNum - minNum + 1) / binNum)
		bin[n] = append(bin[n], arr[j])
		k := len(bin[n]) - 2
		for k >= 0 && arr[j] < bin[n][k] {
			bin[n][k+1] = bin[n][k]
			k--
		}
		bin[n][k+1] = arr[j]
	}

	// 取排序后的数据
	o := 0
	for p, q := range bin {
		for t := 0; t < len(q); t++ {
			arr[o] = bin[p][t]
			o++
		}
	}
	return arr
}

// 基数排序
func RadixSort(arr []int) []int {
	maxNum := arr[0]
	for i := 0; i < len(arr); i++ {
		if maxNum < arr[i] {
			maxNum = arr[i]
		}
	}

	for j := 0; j < len(strconv.Itoa(maxNum)); j++ {
		bin := make([][]int, 10)
		for k := 0; k < len(arr); k++ {
			n := arr[k] / int(math.Pow(10, float64(j))) % 10
			bin[n] = append(bin[n], arr[k])
		}
		m := 0
		for p := 0; p < len(bin); p++ {
			for q := 0; q < len(bin[p]); q++ {
				arr[m] = bin[p][q]
				m++
			}
		}
	}
	return arr
}

// top(k) 问题
func TopK(li []int, k int) []int {
	for i := k/2 - 1; i >= 0; i-- {
		sift(li, i, k-1)
	}
	for j := k; j < len(li); j++ {
		if li[0] < li[j] {
			li[0], li[j] = li[j], li[0]
			sift(li, 0, k-1)
		}
	}
	for n := k - 1; n > 0; n-- {
		li[0], li[n] = li[n], li[0]
		sift(li, 0, n-1)
	}
	return li[:k]
}
func sift(li []int, low, high int) {
	i := low
	j := 2*i + 1
	tmp := li[i]
	for j <= high {
		if j < high && li[j] > li[j+1] {
			j++
		}
		if tmp > li[j] {
			li[i] = li[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	li[i] = tmp
}
