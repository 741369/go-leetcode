/**********************************************
** @Des: meeting-64
** @Author: lzw
** @Last Modified time: 2020/5/23 下午9:51
***********************************************/

package main

func main() {
	println(sumNums(4))
	println(sumNums2(2, 3))
	println(sumNums2(4, 5) / 2)
}

// 递归
func sumNums(n int) int {
	_ = (n > 0) && func() bool {
		n += sumNums(n - 1)
		return true
	}()
	return n
}

/**
int quickMulti(int A, int B) {
    int ans = 0;
    for ( ; B; B >>= 1) {
        if (B & 1) {
            ans += A;
        }
        A <<= 1;
    }
    return ans;
}*/
// 快速乘
func sumNums2(a, b int) int {
	ans := 0
	for ; b > 0; b >>= 1 {
		if b&1 > 0 {
			ans += a
		}
		a <<= 1
	}
	return ans
}
