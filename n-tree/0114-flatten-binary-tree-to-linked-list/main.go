/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/2 下午11:30
***********************************************/

package main

import "fmt"

//给定一个二叉树，原地将它展开为一个单链表。
//
//
//
// 例如，给定二叉树
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6
//
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten(root)

	// 前序打印
	dfs(root)
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

func flatten(root *TreeNode) {
	list := preOrder(root)
	for i := 1; i < len(list); i++ {
		left, right := list[i-1], list[i]
		left.Left, left.Right = nil, right
	}
}

func preOrder(root *TreeNode) []*TreeNode {
	node := make([]*TreeNode, 0)
	if root != nil {
		node = append(node, root)
		node = append(node, preOrder(root.Left)...)
		node = append(node, preOrder(root.Right)...)
	}
	return node
}

//leetcode submit region end(Prohibit modification and deletion)
