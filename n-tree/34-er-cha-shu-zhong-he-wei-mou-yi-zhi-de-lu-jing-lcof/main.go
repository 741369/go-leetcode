/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/12 下午11:07
***********************************************/

package main

import "fmt"

//输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

//
// 示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1

// 返回:
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
//]

// 提示：
//
//
// 节点总数 <= 10000
//
//
// 注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/
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
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	temp := pathSum(root, 22)
	fmt.Println(temp)
}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	recur(root, sum, path, &res)
	return res
}

func recur(root *TreeNode, tar int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	tar -= root.Val
	if root.Left == nil && root.Right == nil && tar == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	recur(root.Left, tar, path, res)
	recur(root.Right, tar, path, res)
	path = path[:len(path)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
