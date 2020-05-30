/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/30 下午5:52
***********************************************/

package main

import (
	"fmt"
	"strings"
)

//以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
//
// 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成
//部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径
//
// 请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表
//示绝对路径的最短字符串。
//
//
//
// 示例 1：
//
// 输入："/home/"
//输出："/home"
//解释：注意，最后一个目录名后面没有斜杠。
//
//
// 示例 2：
//
// 输入："/../"
//输出："/"
//解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级。
//
//
// 示例 3：
//
// 输入："/home//foo/"
//输出："/home/foo"
//解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
//
//
// 示例 4：
//
// 输入："/a/./b/../../c/"
//输出："/c"
//
//
// 示例 5：
//
// 输入："/a/../../b/../c//.//"
//输出："/c"
//
//
// 示例 6：
//
// 输入："/a//b////c/d//././/.."
//输出："/a/b/c"
// Related Topics 栈 字符串
func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}

//leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")
	pathTmp := *new([]string)

	for _, v := range pathArr {
		switch v {
		case ".":
		case "":
		case "..":
			if len(pathTmp) > 0 {
				pathTmp = pathTmp[:len(pathTmp)-1]
			}
		default:
			pathTmp = append(pathTmp, v)
		}
	}
	return "/" + strings.Join(pathTmp, "/")
}

func simplifyPath3(path string) string {
	// 将字符串按照 / 切分
	pathArr := strings.Split(path, "/")
	pathTmp := make([]string, 0) // 模拟栈操作

	// 判断切分后的字符串，如果是 .. 出栈操作，如果不是空字符串，不是 . 就入栈
	for _, v := range pathArr {
		if v == ".." {
			if len(pathTmp) > 0 {
				pathTmp = pathTmp[:len(pathTmp)-1]
			}
		} else if v != "." && v != "" {
			pathTmp = append(pathTmp, v)
		}
	}
	fmt.Println("===========", pathTmp)
	if len(pathTmp) == 0 {
		return "/"
	}
	res := ""
	for _, v := range pathTmp {
		res = res + "/" + v
	}
	return res
}

func simplifyPath2(path string) string {
	// 时间复杂度O(n)，n为path长度
	// 空间复杂度O(n)，n为path长度
	// 需要特殊处理的字符：
	// 1. 多个'/'连续出现：只保留第一个'/'
	// 2. 最后不能是'/'
	// 3. "/./"或位于路径末尾的"/."：只保留第一个'/'
	// 4. "/../"或位于路径末尾的"/.."：去掉之后再往前找到上一个'/'
	// 5. 其余带'.'的情况都是当做名字的一部分来处理
	// 特殊用例：
	// 1. 长度为1，直接返回
	if len(path) == 1 {
		return path
	}
	ret := make([]byte, len(path))
	index := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			// 如果当前'/'是最后一个字符或者前一个字符已经是'/'了，
			// 则当前'/'可以直接抛弃，对应需要特殊处理的字符类型1和类型2
			if i == len(path)-1 || (index > 0 && ret[index-1] == '/') {
				continue
			}
			ret[index] = path[i]
			index++
		} else if path[i] == '.' {
			if (i+1 < len(path) && path[i+1] == '/' && ret[index-1] == '/') ||
				((i == len(path)-1) && ret[index-1] == '/') {
				// 当前位的'.'和下一位的'/'直接抛掉，对应需要特殊处理的字符类型3
				i++
			} else if (i+2 < len(path) && ret[index-1] == '/' && path[i+1] == '.' && path[i+2] == '/') ||
				(i == len(path)-2 && ret[index-1] == '/' && path[i+1] == '.') {
				// "/../"或处于字符串最后的"/.."，对应需要特殊处理的字符类型4
				// index先回到有值的地方
				index--
				for ret[index] != '/' {
					index--
				}
				// 判断是否已经到达了根目录
				if index == 0 {
					index++
				} else {
					// index先回到有值的地方
					index--
					for index > 0 && ret[index] != '/' {
						index--
					}
					// 找到'/'后即完成回退，再指向下一个要写的地方
					index++
				}
				i += 2
			} else {
				// 对应需要特殊处理的字符类型5
				ret[index] = path[i]
				index++
			}
		} else {
			ret[index] = path[i]
			index++
		}
	}
	ret = ret[:index]
	// 若处理完成之后最后一个字符是'/'，则去除最后一个字符，对应需要特殊处理的情况2
	if index != 1 && ret[len(ret)-1] == '/' {
		ret = ret[:len(ret)-1]
	}
	return string(ret)
}

//leetcode submit region end(Prohibit modification and deletion)
