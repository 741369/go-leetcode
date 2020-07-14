/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/6/18 下午11:27
***********************************************/

package main

import "fmt"

//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
// 示例:
//
// 输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
// Related Topics 树 动态规划
func main() {
	fmt.Println(numTrees(3))
}
//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	// 动态规划
	// G(n) = (i=1,n) G(i-1)*G(n-i) G(n)表示长度为n的二叉树的个数
	// F(i, n) 以i为根，序列长度为n的不同二叉搜索树的个数 (1 <= i <= n)
	arr := make([]int, n+1)
	arr[0] = 1 // 空树只有一种情况
	arr[1] = 1 // 单节点树也只有一种情况
	for i := 2; i <= n; i ++ {
		for j := 1; j <= i ; j ++ {
			arr[i] += arr[j-1]*arr[i-j]
		}
	}
	return arr[n]
}
//leetcode submit region end(Prohibit modification and deletion)
