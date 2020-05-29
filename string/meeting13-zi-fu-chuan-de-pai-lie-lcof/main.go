/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/29 下午10:43
***********************************************/

package main

import "fmt"

//输入一个字符串，打印出该字符串中字符的所有排列。
//
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
// 示例:
//
// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
//
// 限制：
//
// 1 <= s 的长度 <= 8
// Related Topics 回溯算法
func main() {
	fmt.Println(permutation("abc"))
	fmt.Println("abc", "===", []byte("abc"))
}

/**
class Solution {
    List<String> res = new LinkedList<>();
    char[] c;
    public String[] permutation(String s) {
        c = s.toCharArray();
        dfs(0);
        return res.toArray(new String[res.size()]);
    }
    void dfs(int x) {
        if(x == c.length - 1) {
            res.add(String.valueOf(c)); // 添加排列方案
            return;
        }
        HashSet<Character> set = new HashSet<>();
        for(int i = x; i < c.length; i++) {
            if(set.contains(c[i])) continue; // 重复，因此剪枝
            set.add(c[i]);
            swap(i, x); // 交换，将 c[i] 固定在第 x 位
            dfs(x + 1); // 开启固定第 x + 1 位字符
            swap(i, x); // 恢复交换
        }
    }
    void swap(int a, int b) {
        char tmp = c[a];
        c[a] = c[b];
        c[b] = tmp;
    }
}

*/
//leetcode submit region begin(Prohibit modification and deletion)
func permutation(s string) []string {
	m := map[string]struct{}{}
	dfs(0, []byte(s), m)
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}

// 深度优先
func dfs(start int, s []byte, m map[string]struct{}) {
	if start == len(s)-1 {
		m[string(s)] = struct{}{} // 保存符合条件的字符
		return
	}
	tmp := map[byte]bool{}
	for i := start; i < len(s); i++ {
		if tmp[s[i]] { // 剪枝
			continue
		}

		tmp[s[i]] = true
		s[i], s[start] = s[start], s[i] // 交换，将 s[i] 固定在第 start 位
		dfs(start+1, s, m)              // 开启固定第start+1位字符
		s[i], s[start] = s[start], s[i] // 恢复交换
	}
}

// 递归 回朔法，效率不是很高
func permutation2(s string) []string {
	m := map[string]struct{}{}
	bfs(s, "", m)

	var res []string
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

// 广度优先
func bfs(remained string, now string, m map[string]struct{}) {
	if len(remained) == 0 {
		m[now] = struct{}{}
	}
	for i := 0; i < len(remained); i++ {
		tmp := remained[:i] + remained[i+1:]
		bfs(tmp, now+string(remained[i]), m)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
