/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/26 下午5:00
***********************************************/

package basic

import "fmt"

type Person struct {
	name string
	sex string
	age int
}

func PointerDemo() {
	var admin *string
	name := "test"
	admin = &name
	fmt.Println(admin)  // 0xc00008e500
	fmt.Println(*admin) // test

	bolden := "tencent"
	admin = &bolden
	fmt.Println(admin)  // 0xc00008e520
	fmt.Println(*admin) // tencent

	major := admin
	*major = "major run"
	fmt.Println(major)  // 0xc000110520
	fmt.Println(bolden) // major run
	fmt.Println(admin == major) // true

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(major) // 0xc000110520
	fmt.Println(charles) // major run
	fmt.Println(bolden) // Charles Bolden

	// 结构体自动解引用
	timmy := &Person{
		name: "Timothy",
		sex:  "man",
		age:  10,
	}
	fmt.Println(timmy)  // &{Timothy man 10}
	fmt.Println(timmy.name) // Timothy
	fmt.Println((*timmy).name) // Timothy

	// 数组自动解引用
	superPowers := &[3]string{"one", "two", "three"}
	fmt.Println(superPowers)  // &[one two three]
	fmt.Println(superPowers[0]) // one
	fmt.Println((*superPowers)[0]) // one
}