/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/7 下午9:38
***********************************************/

package main

import "fmt"
//给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \      \
//        7    2      1
//
//
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	queueVal := []int{root.Val}

	for len(queue) != 0 {
		now := queue[0]
		queue = queue[1:]
		tmp := queueVal[0]
		queueVal = queueVal[1:]

		if now.Left == nil && now.Right == nil {
			if tmp == sum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queue = append(queue, now.Left)
			queueVal = append(queueVal, now.Left.Val + tmp)
		}
		if now.Right != nil {
			queue = append(queue, now.Right)
			queueVal = append(queueVal, now.Right.Val + tmp)
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)

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
		Val:   5,
		Left:  &TreeNode{
			Val:4,
			Left:&TreeNode{
				Val:   11,
				Left:  &TreeNode{Val:7},
				Right: &TreeNode{Val:2},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   13,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: &TreeNode{Val:1},
			},
		},
	}
	fmt.Println(hasPathSum(root, 22))
}

// 递归
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}