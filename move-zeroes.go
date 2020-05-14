/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/14 上午9:35
***********************************************/

package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println("=======", arr)
}

func moveZeroes(arr []int) {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
}
