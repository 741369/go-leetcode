/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午4:47
***********************************************/

package basic

import "fmt"

func TypeDemo()  {
	var a int = 99
	var b float32 = 23.56
	var c = true // 类型推导
	const d byte = 'A'  // 静态类型
	const e string = "test你好" // 编码都采用UTF-8
	f := -1234.567e+78  // 声明变量并赋值自动推导类型

	// 这里的%T，表示按照变量相应值的类型输出
	fmt.Printf("a类型为：%T  值为：%d\n", a, a)
	fmt.Printf("b类型为：%T  值为：%f\n", b, b)
	fmt.Printf("c类型为：%T  值为：%t\n", c, c)
	fmt.Printf("d类型为：%T  值为：%c\n", d, d)
	fmt.Printf("e类型为：%T  值为：%s\n", e, e)
	fmt.Printf("f类型为：%T  值为：%e\n", f, f)

	//a类型为：int  值为：99
	//b类型为：float32  值为：23.559999
	//c类型为：bool  值为：true
	//d类型为：uint8  值为：A
	//e类型为：string  值为：test你好
	//f类型为：float64  值为：-1.234567e+81
}
