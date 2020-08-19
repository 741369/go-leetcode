/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/19 下午8:05
***********************************************/

package main
//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
// 输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
// 输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 输入的字符串长度不会超过 1000 。
//
// Related Topics 字符串 动态规划


//leetcode submit region begin(Prohibit modification and deletion)
// Manacher算法
func countSubstrings(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i ++ {
		t += string(s[i]) + "#"
	}

	n = len(t)
	t += "!"

	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i ++ {
		// 初始化f[i]
		if i <= rMax {
			f[i] = min(rMax - i + 1, f[2*iMax-i])
		} else {
			f[i] = 1
		}

		// 中心扩展
		for t[i+f[i]] == t[i-f[i]] {
			f[i] ++
		}

		// 动态维护
		if i + f[i] - 1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}

		// 统计答案  (f[i] - 1) /2 上取整
		ans += f[i] / 2
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//leetcode submit region end(Prohibit modification and deletion)

// 中心扩展
func countSubstrings2(s string) int {
	// 遍历字符串的所有中心点，中心点可能一个数，也可能2个数
	ans, n := 0, len(s)
	for i := 0; i < 2 * n - 1; i ++ {
		left, right := i / 2, i / 2 + i % 2
		for left >= 0 && right < n && s[left] == s[right] {
			left --
			right ++
			ans ++
		}
	}
	return ans
}
