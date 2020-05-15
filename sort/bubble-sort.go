/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/11 下午9:43
***********************************************/

package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
)

func main() {
	arr := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("冒泡====", BubbleSort(arr))

	arr2 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("插入====", InsertSort(arr2))

	arr3 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("选择====", SelectionSort(arr3))

	arr4 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("希尔====", ShellSort(arr4))

	arr5 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("快排====", QuickSort(arr5))
	fmt.Println("快排====", QuickSort(nums))

	arr52 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("快排2====", QuickSort2(arr52))

	arr6 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("堆排====", HeapSort(arr6))

	arr7 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("归并====", MergeSort(arr7))

	arr8 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("计数====", CountSort(arr8))

	arr9 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("桶排序====", BinSort(arr9))

	arr10 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("基数====", RadixSort(arr10))

	arr11 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("Top(n)====", TopK(arr11, 5))

	arr12 := []int{47, 29, 71, 78, 19, 24, 47}
	//arr12 := []int{3, 2, 4, 1, 6, 5, 7, 9, 8}
	fmt.Println("Top(n)====", TopK2(arr12, 6))
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
		p := partition(arr, left, right)
		quickSort(arr, left, p-1)
		quickSort(arr, p+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	i, j, temp := left, right, arr[left]
	for i < j {
		for i < j && temp < arr[j] {
			j--
		}
		for i < j && temp >= arr[i] {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[left] = arr[left], arr[i]
	return i
}

// 快速排序2
func QuickSort2(arr []int) []int {
	quickSort2(arr, 0, len(arr)-1)
	return arr
}
func quickSort2(arr []int, left int, right int) {
	if left < right {
		p := partition2(arr, left, right)
		quickSort2(arr, left, p-1)
		quickSort2(arr, p+1, right)
	}
}

func partition2(arr []int, left int, right int) int {
	i, j, temp := left, left, arr[right]
	for ; j < right; j++ {
		if arr[j] < temp {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[right], arr[i] = arr[i], arr[right]
	return i
}

// 堆排序
func HeapSort(arr []int) []int {
	for i := len(arr) / 2; i >= 0; i-- {
		initHeap(arr, i, len(arr))
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		initHeap(arr, 0, i)
	}
	return arr
}

/**
 *  下沉操作，执行删除操作相当于把最后
 *  * 一个元素赋给根元素之后，然后对根元素执行下沉操作
 * @param arr
 * @param parent 要下沉元素的下标
 * @param length 数组长度
 */
func initHeap(nums []int, parent, len int) {
	//临时保证要下沉的元素
	temp := nums[parent]
	//定位左孩子节点位置
	child := 2*parent + 1

	//开始下沉
	for child < len {
		//如果右孩子节点比左孩子小，则定位到右孩子
		if child+1 < len && nums[child] < nums[child+1] {
			child++
		}
		//如果父节点比孩子节点小或等于，则下沉结束
		if child < len && nums[child] <= temp {
			break
		}
		//单向赋值
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
	i, j := left, mid+1
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
/**
a.先取前k个数建小根堆，这样就能保证堆顶元素是整个堆的最小值;
b.然后遍历列表的k到最后，如果值比堆顶大，就和堆顶交换，交换完后再对堆建小根堆，这样就能保证交换完后，堆顶元素仍然是整个堆的最小值；
c.一直遍历到列表的最后一个值，这一步做完之后，就保证了整个列表最大的前k个数已经放进了堆里；
d.按数的大到小出堆；
*/
func TopK(arr []int, k int) []int {
	for i := k/2 - 1; i >= 0; i-- {
		sift(arr, i, k-1)
	}
	for j := k; j < len(arr); j++ {
		if arr[0] < arr[j] {
			arr[0], arr[j] = arr[j], arr[0]
			sift(arr, 0, k-1)
		}
	}
	for n := k - 1; n > 0; n-- {
		arr[0], arr[n] = arr[n], arr[0]
		sift(arr, 0, n-1)
	}
	return arr[:k]
}
func sift(arr []int, left, right int) {
	i, j := left, 2*left+1
	tmp := arr[left]
	for j <= right {
		if j < right && arr[j] > arr[j+1] {
			j++
		}
		if tmp > arr[j] {
			arr[i] = arr[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	arr[i] = tmp
}

func TopK2(arr []int, k int) []int {
	TopK2Sort(arr, 0, len(arr)-1, k)
	return arr[len(arr)-k:]
}

func TopK2Sort(arr []int, left, right, k int) {
	if left < right {
		p := partition(arr, left, right)
		if p == k-1 {
			return
		} else if p > k-1 {
			TopK2Sort(arr, left, p-1, k)
		} else {
			TopK2Sort(arr, p+1, right, k)
		}
	}
}

//堆排序(大顶堆)
func smallestK(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 || k == 0 {
		return nil
	}
	res := make([]int, 0)
	//构建堆
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(arr); i++ {
		if h.Len() < k { //维持一个k大小的堆
			heap.Push(h, arr[i])
		} else { //当新加入数时，如果当前堆顶
			//元素大于当前数，把堆顶元素弹出，加入新数
			top := heap.Pop(h)
			if top.(int) > arr[i] {
				heap.Push(h, arr[i])
			} else {
				heap.Push(h, top)
			}
		}
	}
	for h.Len() != 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

//构建大顶堆
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
