/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/19 上午12:04
***********************************************/

package main

import "fmt"

//二叉搜索树中的两个节点被错误地交换。
//
// 请在不改变其结构的情况下，恢复这棵树。
//
// 示例 1:
//
// 输入: [1,3,null,null,2]
//
//   1
//  /
// 3
//  \
//   2
//
//输出: [3,1,null,null,2]
//
//   3
//  /
// 1
//  \
//   2
//
//
// 示例 2:
//
// 输入: [3,1,4,null,null,2]
//
//  3
// / \
//1   4
//   /
//  2
//
//输出: [2,1,4,null,null,3]
//
//  2
// / \
//1   4
//   /
//  3
//
// 进阶:
//
//
// 使用 O(n) 空间复杂度的解法很容易实现。
// 你能想出一个只使用常数空间的解决方案吗？
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
	1. 找到二叉搜索树中序遍历得到值序列的不满足条件的位置。
	2. 如果有2个节点i，j满足，i < j && a[i] > a[i+1] && a[j] > a[j+1],则交换i, j+1即可
	3. 如果只有1个节点i，满足，a[i] > a[i+1]，则交换i 和 i+1即可
 */
func recoverTree(root *TreeNode)  {
	var nums []int

	// 中序遍历二叉树
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)
	x, y := findTwoSwapped(nums)
	fmt.Println(x, y)
	recover(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i + 1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Right, count, x, y)
	recover(root.Left, count, x, y)
}

//leetcode submit region end(Prohibit modification and deletion)
