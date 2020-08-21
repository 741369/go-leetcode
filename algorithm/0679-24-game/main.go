/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/8/22 上午12:54
***********************************************/

package main

import "fmt"

//你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。
//
// 示例 1:
//
// 输入: [4, 1, 8, 7]
//输出: True
//解释: (8-4) * (7-1) = 24
//
//
// 示例 2:
//
// 输入: [1, 2, 1, 2]
//输出: False
//
//
// 注意:
//
//
// 除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
// 每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允
//许的。
// 你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。
//
// Related Topics 深度优先搜索
func main() {
	fmt.Println(judgePoint24([]int{4,1,8,7}))
	fmt.Println(judgePoint24([]int{1,2,1,2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
const (
	TARGET = 24
	EPSILON = 1e-6
	ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)
func judgePoint24(nums []int) bool {
	var list []float64
	for _, num := range nums {
		list = append(list, float64(num))
	}

	return solve(list)
}

func solve(list []float64) bool {
	if len(list) == 0 {
		return false
	}
	if len(list) == 1 {
		return abs(list[0] - TARGET) < EPSILON
	}
	size := len(list)
	for i := 0; i < size; i ++ {
		for j := 0; j < size; j ++ {
			if i != j {
				list2 := []float64{}
				for k := 0; k < size; k++ {
					if k != i && k != j {
						list2 = append(list2, list[k])
					}
				}
				for k := 0; k < 4; k++ {
					if k < 2 && i < j {
						continue
					}
					switch k {
					case ADD:
						list2 = append(list2, list[i] + list[j])
					case MULTIPLY:
						list2 = append(list2, list[i] * list[j])
					case SUBTRACT:
						list2 = append(list2, list[i] - list[j])
					case DIVIDE:
						if abs(list[j]) < EPSILON {
							continue
						} else {
							list2 = append(list2, list[i] / list[j])
						}
					}
					if solve(list2) {
						return true
					}
					list2 = list2[:len(list2) - 1]
				}
			}
		}
	}
	return false
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)

