/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/15 上午12:36
***********************************************/

package main

//给定一个二叉树，返回它的 前序 遍历。
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
//输出: [1,2,3]
//
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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代 先序
func preorderTraversal(root *TreeNode) []int {
	var (
		res   []int
		stack []*TreeNode
	)

	for len(stack) > 0 || root != nil {
		// 最左子树入栈
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		root = stack[index].Right
		stack = stack[:index]
	}
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			stack = append(stack, cur)
			stack = append(stack, nil)
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, tmp.Val)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 递归先序
func preorderdfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderdfs(root.Left)...)
	res = append(res, preorderdfs(root.Right)...)

	return res
}
