/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/15 上午12:33
***********************************************/

package main

//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树


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
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val}, res...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return res
}

// 迭代 后序2
func postorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			stack = append(stack, cur)
			stack = append(stack, nil)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, tmp.Val)
		}
	}
	return res
}

// 递归后序
func postorderdfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, postorderdfs(root.Left)...)
	res = append(res, postorderdfs(root.Right)...)
	res = append(res, root.Val)
	return res
}