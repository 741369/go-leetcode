/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/5 下午4:35
***********************************************/

package main

import "fmt"

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
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
func main() {
	//
	node := &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{Val:3},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{Val:1},
		},
	}
	fmt.Println(rob(node))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 动态规划方法
func rob(root *TreeNode) int {
	// f(o) 表示节点o被选中，o节点的子树上被选择的节点的最大权值和，
	// g(o) 表示节点o不被选中，o节点的子树上被选择的节点的最大权值和
	// l,r分别为o的左右子树
	// o被选中，则左右子树不会被选中，f(o) = g(l) + g(r)
	// o不被选中，则左右子树都可以被选中，g(o) = max(f(l), g(l)) + max(f(r), g(r))
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	// 数组第一个值f, 第二个值g
	left, right := dfs(node.Left), dfs(node.Right)
	selected := node.Val + left[1] + right[1]
	notSelected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{selected, notSelected}

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//leetcode submit region end(Prohibit modification and deletion)

