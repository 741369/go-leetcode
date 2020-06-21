/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/21 下午11:27
***********************************************/

package main

import (
	"fmt"
	"math"
)

//给定一个非空二叉树，返回其最大路径和。
//
// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
// 示例 1:
//
// 输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//
//
// 示例 2:
//
// 输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42
// Related Topics 树 深度优先搜索

// -1
// /  \
// -2    3
// /  \
// 4   6
//      \
//       7

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
	root := &TreeNode{
		Val:   -1,
		Left:  &TreeNode{
			Val:-2,
			Left:&TreeNode{Val:4},
			Right:&TreeNode{
				Val:   6,
				Right: &TreeNode{Val:7},
			},
		},
		Right: &TreeNode{
			Val:   3,
		},
	}
	fmt.Println(maxPathSum(root))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxVal := math.MinInt64
	var maxGain func (node *TreeNode) int

	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 该节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值，
		// 左子树和右子树的最大路径和肯定> 0 ，所以不用考虑负数
		currentNodePath := node.Val + leftGain + rightGain

		// 更新最大值
		maxVal = max(maxVal, currentNodePath)

		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxVal
}
//leetcode submit region end(Prohibit modification and deletion)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}