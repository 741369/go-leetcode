/**********************************************
** @Des: 
** @Author: liuzhiwang
** @Last Modified time: 2020/6/3 下午6:23
***********************************************/

package main

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

func main()  {
	for i := 0; i < 10; i ++ {
		println(t(int64(i)))
	}
}

func t(i int64) int64 {
	node, err := snowflake.NewNode(i)
	if err != nil {
		return 0
	}
	fmt.Println(node.Generate().Base2())
	fmt.Println(string(node.Generate().Bytes()))
	fmt.Println(node.Generate().Base36())

	return node.Generate().Int64()
}
