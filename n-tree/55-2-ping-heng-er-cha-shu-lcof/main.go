/**********************************************
** @Des:
** @Author: liuzhiwang
** @Last Modified time: 2020/6/13 上午12:08
***********************************************/

package main

/**

输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。
*/

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
class Solution {
    public boolean isBalanced(TreeNode root) {
        return recur(root) != -1;
    }

    private int recur(TreeNode root) {
        if (root == null) return 0;
        int left = recur(root.left);
        if(left == -1) return -1;
        int right = recur(root.right);
        if(right == -1) return -1;
        return Math.abs(left - right) < 2 ? Math.max(left, right) + 1 : -1;
    }
}

*/
func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	if abs(left, right) < 2 {
		return max(left, right) + 1
	}
	return -1
}

func abs(left, right int) int {
	tmp := left - right
	if tmp > 0 {
		return tmp
	}
	return -tmp
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
