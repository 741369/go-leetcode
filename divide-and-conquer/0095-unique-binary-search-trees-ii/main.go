/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/25 上午9:44
***********************************************/
//给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
//
// 示例:
//
// 输入: 3
//输出:
//[
//  [1,null,3,2],
//  [3,2,null,1],
//  [3,1,null,null,2],
//  [2,1,3],
//  [1,null,2,null,3]
//]
//解释:
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
//
// Related Topics 树 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	trees := generateTrees(4)
	fmt.Println("==========", trees)
	for k, v := range trees {
		fmt.Println("=========", k, "===", v)
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(left, right int) (res []*TreeNode) {
	if left > right {
		res = append(res, nil)
		return
	}
	for i := left; i <= right; i++ {
		leftTree := generate(left, i-1)
		rightTree := generate(i+1, right)
		for _, leftTmp := range leftTree {
			for _, rightTmp := range rightTree {
				tmp := &TreeNode{
					Val:   i,
					Left:  leftTmp,
					Right: rightTmp,
				}
				res = append(res, tmp)
			}
		}
	}
	return res
}
