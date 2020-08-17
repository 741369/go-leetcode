/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/17 下午12:19
***********************************************/

package main

import "fmt"

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//
// 示例 1:
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回 true 。
//
//示例 2:
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//        1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//
//
// 返回 false 。
//
//
// Related Topics 树 深度优先搜索


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func main() {
	node := &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isBalanced(node))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(maxHeight(root.Left), maxHeight(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxHeight(node.Left), maxHeight(node.Right)) + 1
}

func max(i, j int) int {
	if i >j {
		return i
	}
	return j
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
//leetcode submit region end(Prohibit modification and deletion)

