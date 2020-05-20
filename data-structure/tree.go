/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/20 下午1:01
***********************************************/

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// [1,null,2,3]
func main() {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	arr := preorderTraversal(tree)
	fmt.Println("===========", arr)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var rootTemp = root
	var stack []*TreeNode
	var arr []int

	for rootTemp != nil || len(stack) > 0 {
		for rootTemp != nil {
			arr = append(arr, rootTemp.Val)
			stack = append(stack, rootTemp.Right)
			rootTemp = rootTemp.Left
		}

		rootTemp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return arr
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var arr []int
	arr = append(arr, root.Val)
	arr = append(arr, preorderTraversal2(root.Left)...)
	arr = append(arr, preorderTraversal2(root.Right)...)
	return arr
}
