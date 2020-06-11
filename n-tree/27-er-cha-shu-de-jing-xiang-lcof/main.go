/**********************************************
** @Des: 0026 https://leetcode-cn.com/problems/invert-binary-tree/
** @Author: lzw
** @Last Modified time: 2020/6/10 下午10:14
***********************************************/

package main

import "fmt"

//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
// 例如输入：
//
// 4
// / \
// 2 7
// / \ / \
//1 3 6 9
//镜像输出：
//
// 4
// / \
// 7 2
// / \ / \
//9 6 3 1
//
//
//
// 示例 1：
//
// 输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 1000
//
// 注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/
// Related Topics 树

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
	/*root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}*/
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	dfs(root)
	temp := mirrorTree(root)
	fmt.Println("")
	dfs(temp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
方法一：递归法
根据二叉树镜像的定义，考虑递归遍历（dfs）二叉树，交换每个节点的左 / 右子节点，即可生成二叉树的镜像。
*/
func mirrorTree(root *TreeNode) *TreeNode {
	// 当节点 rootroot 为空时（即越过叶节点），则返回 nullnull ；
	if root == nil {
		return nil
	}
	// 暂存root的left节点，因为root.right遍历完后会覆盖root.left
	temp := root.Left
	// 开始递归右子树，并将返回值赋值给左子树，
	root.Left = mirrorTree(root.Right)
	// 开始递归保存的左子树，并将结果赋值给右子树
	root.Right = mirrorTree(temp)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
/**
方法二：使用栈或者队列辅助
特例处理： 当 root 为空时，直接返回 null ；
初始化： 栈（或队列），本文用栈，并加入根节点 root 。
循环交换： 当栈 stack 为空时跳出；
出栈： 记为 node ；
添加子节点： 将 node 左和右子节点入栈；
交换： 交换 node 的左 / 右子节点。
*/
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		node.Right, node.Left = node.Left, node.Right
	}
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d\t", root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
