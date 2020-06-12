/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/12 下午10:52
***********************************************/

package main

import "fmt"

//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
//
//
//
// 例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
//
//
//
//
// 提示：
//
//
// 节点总数 <= 1000
//
// Related Topics 树 广度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	temp := levelOrder(root)
	fmt.Println(temp)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	var res [][]int
	flag := true
	stack = append(stack, root)
	for len(stack) > 0 {
		var tmp []int
		for i := len(stack) - 1; i >= 0; i-- {
			node := stack[0]
			stack = stack[1:]
			if flag {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		flag = !flag
		res = append(res, tmp)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
