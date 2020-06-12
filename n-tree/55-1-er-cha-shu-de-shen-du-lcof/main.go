/**********************************************
** @Des:
** @Author: liuzhiwang
** @Last Modified time: 2020/6/12 下午11:41
***********************************************/

package main

import "fmt"

/**
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

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

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
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

	temp := maxDepth(root)
	fmt.Println(temp)
}

func maxDepth(root *TreeNode) int {
	var path []int
	return dfs(root, path, 0)
}

func dfs(root *TreeNode, path []int, depth int) int {
	if root == nil {
		return depth
	}

	path = append(path, root.Val)
	fmt.Printf("====%v, %v\n", path, root)
	if root.Left == nil && root.Right == nil && len(path) > depth {
		depth = len(path)
	}
	depth = dfs(root.Left, path, depth)
	depth = dfs(root.Right, path, depth)
	path = path[:len(path)-1]
	return depth
}
