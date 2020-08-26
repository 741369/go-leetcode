/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午5:00
***********************************************/

package basic

import "fmt"

// 切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度(len)和切片的容量(cap)

func SliceDemo()  {
	a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := a[3:6] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// s:[3 4 5] len(s):3 cap(s):5

	// 切片修改
	s[1] = 111
	fmt.Println(a)
	// [0 1 2 3 111 5 6 7]

	fmt.Println(s)
	// [3 111 5]

	// 拷贝切片
	c := make([]int, 3, 5)  // 定义一个切片，指定长度和容量
	copy(c, s) // 使用copy()函数将切片s中的元素复制到切片c中
	c[1] = 2222

	fmt.Println(c)
	// [3 2222 5]

	fmt.Println(s)
	// [3 111 5]
}