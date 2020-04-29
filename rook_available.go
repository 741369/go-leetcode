/**********************************************
** @Des: 999. available-captures-for-rook
** @Author: liuzhiwang
** @Last Modified time: 2020/3/26 下午5:45
***********************************************/

package main

import "fmt"

func main() {
	/**
	[['.','.','.','.','.','.','.','.'],['.','.','.','p','.','.','.','.'],['.','.','.','R','.','.','.','p'],['.','.','.','.','.','.','.','.'],['.','.','.','.','.','.','.','.'],['.','.','.','p','.','.','.','.'],['.','.','.','.','.','.','.','.'],['.','.','.','.','.','.','.','.']]
	*/
	//str := map[string]map[string]string[
	//

	str := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	num := numRookCaptures(str)
	fmt.Println("===", num)

}

func numRookCaptures(board [][]byte) int {
	boardLength := 8
	var row, col, ret int
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}

	for i := 0; i < boardLength; i++ {
		flag := false
		for j := 0; j < boardLength; j++ {
			if board[i][j] == 'R' {
				row = i
				col = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}

	for i := 0; i < 4; i++ {
		for step := 0; true; step++ {
			tx := row + step*dx[i]
			ty := col + step*dy[i]
			if tx < 0 || tx >= 8 || ty < 0 || ty >= 8 || board[tx][ty] == 'B' {
				break
			}
			if board[tx][ty] == 'p' {
				ret++
				break
			}
		}
	}

	/*for i := col + 1; i < boardLength; i++ {
		if board[row][i] == 'p' {
			ret++
			break
		} else if board[row][i] == 'B' {
			break
		}
	}
	for i := col - 1; i >= 0; i-- {
		if board[row][i] == 'p' {
			ret++
			break
		} else if board[row][i] == 'B' {
			break
		}
	}
	for i := row + 1; i < boardLength; i++ {
		if board[i][col] == 'p' {
			ret++
			break
		} else if board[i][col] == 'B' {
			break
		}
	}
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == 'p' {
			ret++
			break
		} else if board[i][col] == 'B' {
			break
		}
	}
	*/
	return ret
}
