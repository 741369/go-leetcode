package main

import "fmt"
// 95. 不同的二叉搜索树 II
// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。

 

// 示例：

// 输入：3
// 输出：
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// 解释：
// 以上的输出对应以下 5 种不同结构的二叉搜索树：

//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3
 

// 提示：

// 0 <= n <= 8

func main(){
	fmt.Println(generateTrees(3))
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
 func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }
    return helperGenerate(1, n)
}

func helperGenerate(left, right int) ([]*TreeNode) {
    if left > right {
        return []*TreeNode{nil}
    }
    res := []*TreeNode{}
    for i := left; i <= right; i ++ {
        leftNode := helperGenerate(left, i-1)
        rightNode := helperGenerate(i+1, right)
        for _, v1 := range leftNode {
            for _, v2 := range rightNode {
                tmp := &TreeNode{Val: i}
                tmp.Left = v1
                tmp.Right = v2
                res = append(res, tmp) 
            }
        }
    }
    return res
}