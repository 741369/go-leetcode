/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/14 下午10:31
***********************************************/

package main

import "fmt"

//给定一个二叉树，返回它的中序 遍历。
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
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表


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

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left:&TreeNode{
				Val:   4,
			},
			Right:&TreeNode{
				Val:   5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Printf("preorder迭代 =====  %#v\n", preorderTraversal(root))
	fmt.Printf("inorder迭代 =====  %#v\n", inorderTraversal(root))
	fmt.Printf("postorder迭代 =====  %#v\n", postorderTraversal(root))

	fmt.Printf("preorder递归 ===== %#v\n", preorderdfs(root))
	fmt.Printf("inorder递归 ===== %#v\n", inorderdfs(root))
	fmt.Printf("postorder递归 ===== %#v\n", postorderdfs(root))
}

// 迭代 先序
func preorderTraversal(root *TreeNode) []int {
	var (
		res []int
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

// 迭代 中序
func inorderTraversal(root *TreeNode) []int {
	var (
		res []int
		stack []*TreeNode
	)

	for len(stack) > 0 || root != nil {
		// 最左子树入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		res = append(res, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}
	return res
}

// 迭代 后序
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

// 递归
func inorderTraversal2(root *TreeNode) []int {
	return inorderdfs(root)
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

// 递归中序
func inorderdfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderdfs(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderdfs(root.Right)...)
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