/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/7/3 下午11:20
***********************************************/

package main

import "fmt"

//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
// Related Topics 树 深度优先搜索


func main() {
	node := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	printTree(node)
}
func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	printTree(node.Left)
	fmt.Println(node.Val)
	printTree(node.Right)
}

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

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}
//leetcode submit region end(Prohibit modification and deletion)

func dfs(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (right + left) /2
	return &TreeNode{
		Val:   nums[mid],
		Left:  dfs(nums, left, mid - 1),
		Right: dfs(nums, mid+1, right),
	}
}