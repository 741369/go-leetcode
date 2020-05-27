package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 1, 1, 2, 0, 0}
	fmt.Println("==", sortArray(arr))
}

func quicksort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	index := partition(nums, left, right)

	quicksort(nums, left, index-1)
	quicksort(nums, index+1, right)
}

func partition(nums []int, left int, right int) int {
	i := left
	j := right
	for i != j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}

func sortArray6(nums []int) []int {
	// 从n/2  最后一个非叶子结点起开始构建大顶堆
	for i := len(nums) / 2; i >= 0; i-- {
		heapSort(nums, i)
	}

	end := len(nums) - 1
	// 每次将大顶堆的最大值与末尾进行交换，并再次排序
	for end > 0 {
		nums[0], nums[end] = nums[end], nums[0]
		heapSort(nums[:end], 0)
		end--
	}
	return nums

}

// 对一个非叶子结点进行排序
func heapSort(nums []int, pos int) {
	end := len(nums) - 1
	left := 2*pos + 1

	if left > end {
		return
	}

	right := 2*pos + 2
	temp := left

	// 先左右子结点进行比较，找出较小的那一个
	if right <= end && nums[right] > nums[temp] {
		temp = right
	}

	if nums[temp] <= nums[pos] {
		return
	}

	nums[temp], nums[pos] = nums[pos], nums[temp]

	// 如果发生了交换的话 就要继续调查后续子节点(只调查交换了的后续，不用全调查，不然会超时)
	heapSort(nums, temp)
}

// 归并排序
func sortArray5(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	partA := sortArray(nums[:len(nums)/2])
	partB := sortArray(nums[len(nums)/2:])

	temp := make([]int, len(partA)+len(partB))

	aPointer := 0
	bPointer := 0
	i := 0

	for aPointer < len(partA) && bPointer < len(partB) {
		if partA[aPointer] < partB[bPointer] {
			temp[i] = partA[aPointer]
			aPointer++
		} else {
			temp[i] = partB[bPointer]
			bPointer++
		}
		i++
	}
	for aPointer < len(partA) {
		temp[i] = partA[aPointer]
		aPointer++
		i++
	}
	for bPointer < len(partB) {
		temp[i] = partB[bPointer]
		bPointer++
		i++
	}
	return temp
}

func sortArray4(nums []int) []int {
	quickSort(nums)
	return nums
}

// 快速排序
func quickSort(nums []int) {
	left, right := 0, len(nums)-1
	for right > left {
		// 右边部分放大于
		if nums[right] > nums[0] {
			right--
			continue
		}
		// 左边部分放小于等于
		if nums[left] <= nums[0] {
			left++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[0], nums[right] = nums[right], nums[0]
	if len(nums[:right]) > 1 {
		sortArray(nums[:right])
	}
	if len(nums[right+1:]) > 1 {
		sortArray(nums[right+1:])
	}
}

// 慢
func sortArray3(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		flag := true
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return nums
}

/**
func sortArray(nums []int) []int {
    // 冒泡排序
    // 两个两个比较大小，将较大值放到后面
    for i := 0; i < len(nums) - 1; i++ {
        flag := true
        for j := 0; j < len(nums) - 1 - i; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
                flag = false
            }
        }
        if flag {// 从头到尾都没有进行交换，说明已经排序完成
            break
        }
    }
    return nums
}
*/

// 交换排序
func sortArray2(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}
