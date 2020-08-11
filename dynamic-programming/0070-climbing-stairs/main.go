/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/10 下午9:43
***********************************************/

package main

import "fmt"

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划
func main() {
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
	fmt.Println(climbStairs2(20))
}

//leetcode submit region begin(Prohibit modification and deletion)
// 动态规划
func climbStairs(n int) int {
	first, second := 0, 1
	for i := 1; i <= n; i ++ {
		first += second
		first, second = second, first
	}
	return second
}

func climbStairs5(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first += second
		first, second = second, first
	}
	return second
}

func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i <= n; i ++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}

// 递归
func climbStairs3(n int) int {
    // 1: 1
    // 2: 2
    // 3: f(1) + f(2), mutual exclusive,complete exhaustive
    // 4: f(2) + f(3)
    // f(n) = f(n-1) + f(n-2). fibonacci
    if n <= 2 {
        return n
    }
    return climbStairs3(n-2) + climbStairs3(n-1)
}
//leetcode submit region end(Prohibit modification and deletion)
