/**********************************************
** @Des: 70. 爬楼梯
** https://leetcode-cn.com/problems/climbing-stairs/
** 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
** @Author: lzw
** @Last Modified time: 2020/5/15 上午9:07
***********************************************/

package main

import "fmt"

func main() {
	fmt.Println("====", climbStairsDynamic(44))
}

func climbStairs(n int) int {
	arr := make([]int, n+1)
	return climbStairsArr(n, arr)
}

/**
 * @Author liu
 * @Description // 动态规划方法
 * @Date 2020/5/15 上午10:01
 * @Param
 * @return
 **/
func climbStairsDynamic(n int) int {
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

/**
 * @Author liu
 * @Description // 爬楼梯记忆法
 * @Date 2020/5/15 上午9:37
 * @Param
 * @return
 **/
func climbStairsArr(n int, arr []int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	if arr[n] > 0 {
		return arr[n]
	}
	arr[n] = climbStairsArr(n-1, arr) + climbStairsArr(n-2, arr)
	return arr[n]
}

/**
 * @Author liu
 * @Description // 蛮力递归法
 * 时间复杂度 O(2^n)
 * 空间复杂度 O(n)
 * @Date 2020/5/15 上午9:34
 * @Param 楼梯阶数 n
 * @return 返回多少种爬法
 **/
func climbStairs2(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}
