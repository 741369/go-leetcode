//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回锯齿形层次遍历如下：
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := zigzagLevelOrder(tree)
	fmt.Printf("%#v\n", res)
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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs搜索
func zigzagLevelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	flag := true
	for len(q) > 0 {
		qLen := len(q)
		ans := make([]int, qLen)
		for i := 0; i < qLen; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if flag {
				ans[i] = node.Val
			} else {
				ans[qLen-i-1] = node.Val
			}
		}
		res = append(res, ans)
		flag = !flag
		q = q[qLen:]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// dfs搜索

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if level == len(*res) {
		*res = append(*res, []int{})
	}
	// } else {
	if level%2 == 0 {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		(*res)[level] = append([]int{root.Val}, (*res)[level]...)
	}
	// }
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)
}
