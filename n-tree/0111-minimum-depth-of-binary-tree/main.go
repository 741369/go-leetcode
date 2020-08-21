/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/21 下午7:00
***********************************************/

package main

import (
	"fmt"
	"math"
)

//给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//
// 给定二叉树 [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最小深度 2.
// Related Topics 树 深度优先搜索 广度优先搜索


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
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(minDepth(root))

	root2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 3},
		Right: nil,
	}
	fmt.Println(minDepth(root2))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minTmp := math.MaxInt64
	if root.Left != nil {
		minTmp = min(minDepth(root.Left), minTmp)
	}
	if root.Right != nil {
		minTmp = min(minDepth(root.Right), minTmp)
	}
	return minTmp + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
//leetcode submit region end(Prohibit modification and deletion)
