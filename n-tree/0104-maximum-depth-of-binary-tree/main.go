package main

// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。

// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3 。

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //   迭代，棧
 func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{root}
    ans := 0
    for len(queue) > 0 {
        sz := len(queue)
        for sz > 0 {
            tmp := queue[0]
            queue = queue[1:]
            if tmp.Left != nil {
                queue = append(queue, tmp.Left)
            }
            if tmp.Right != nil {
                queue = append(queue, tmp.Right)
            }
            sz --
        }
        ans ++
    }

    return ans
}

// 递归
func maxDepth2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}