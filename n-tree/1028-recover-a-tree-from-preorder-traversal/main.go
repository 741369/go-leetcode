/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/18 上午12:09
***********************************************/

package main

import "fmt"

//我们从二叉树的根节点 root 开始进行深度优先搜索。
//
// 在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。
//根节点的深度为 0）。
//
// 如果节点只有一个子节点，那么保证该子节点为左子节点。
//
// 给出遍历输出 S，还原树并返回其根节点 root。
//
//
//
// 示例 1：
//
//
//
// 输入："1-2--3--4-5--6--7"
//输出：[1,2,5,3,4,6,7]
//
//
// 示例 2：
//
//
//
// 输入："1-2--3---4-5--6---7"
//输出：[1,2,5,3,null,6,null,4,null,7]
//
//
// 示例 3：
//
//
//
// 输入："1-401--349---90--88"
//输出：[1,401,null,349,88,90]
//
//
//
//
// 提示：
//
//
// 原始树中的节点数介于 1 和 1000 之间。
// 每个节点的值介于 1 和 10 ^ 9 之间。
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
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node := recoverFromPreorder("1-2--3--4-5--6--7")
	dfs(node)
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d\t", root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

func recoverFromPreorder(S string) *TreeNode {
	node, pos := []*TreeNode{}, 0

	for pos < len(S) {
		// 先找当前几点所在层级
		level := 0
		for S[pos] == '-' {
			pos ++
			level ++
		}

		// 获取当前节点的值
		value := 0
		for ;pos < len(S) && S[pos] >= '0' && S[pos] <= '9'; pos ++ {
			value = value * 10 + int(S[pos] - '0')
		}

		tmp := &TreeNode{Val:value}
		// 判断当前节点是左子节点还是右子节点
		if level == len(node) { // 判断节点是否为左子树
			if len(node) > 0 {
				node[level-1].Left = tmp
			}
		} else {
			node = node[:level]
			node[level-1].Right = tmp
		}
		node = append(node, tmp)
	}

	return node[0]
}
//leetcode submit region end(Prohibit modification and deletion)

