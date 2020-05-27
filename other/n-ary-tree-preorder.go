/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/16 下午9:47
***********************************************/

package main

import "fmt"

func main() {
	//-----------多叉树层次遍历（打印行号）-----------------
	node11 := &Node{Val: 10}
	node10 := &Node{Val: 10}
	node9 := &Node{Val: 9}
	node8 := &Node{Val: 8}
	node7 := &Node{Val: 7}
	node6 := &Node{Val: 6}
	node5 := &Node{Val: 5, Children: []*Node{node8, node9, node10, node11}}
	node4 := &Node{Val: 4}
	node3 := &Node{Val: 3}
	node2 := &Node{Val: 2, Children: []*Node{node5, node6, node7}}
	root := &Node{Val: 1, Children: []*Node{node2, node3, node4}}
	fmt.Println("=========", preorder(root))
}

type Node struct {
	Val      int
	Children []*Node
}

/**
 * @Author liu
 * @Description // TODO 迭代解法
 * @Date 2020/5/16 下午10:04
 * @Param
 * @return
 **/
func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val) //前序输出
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i]) //入栈
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

/**
 * @Author liu
 * @Description // TODO递归解法
 * @Date 2020/5/16 下午10:03
 * @Param
 * @return
 **/
var res []int

func preorder2(root *Node) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *Node) {
	if root != nil {
		res = append(res, root.Val)
	}
	for _, v := range root.Children {
		dfs(v)
	}
}
