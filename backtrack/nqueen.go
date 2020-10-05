package backtrack

import (
	"bytes"
	"fmt"
)

type checkBoard [][]rune

func (cb checkBoard) String() string {

	var buffer bytes.Buffer
	buffer.WriteRune('\n')
	for row := range cb {
		buffer.WriteString(fmt.Sprintf("%s\n", string(cb[row])))
	}

	return buffer.String()
}

func (cb checkBoard) Copy() checkBoard {
	nb := makeCheckBoard(len(cb))
	for row := range nb {
		copy(nb[row], cb[row])
	}
	return nb
}

func makeCheckBoard(n int) checkBoard {

	var board checkBoard
	for i := 0; i < n; i++ {
		r := make([]rune, n)
		for j := 0; j < n; j++ {
			r[j] = '.'
		}
		board = append(board, r)
	}

	return board
}

func NQueen(n int) []checkBoard {

	cb := makeCheckBoard(n)

	var res []checkBoard

	// isvalid function, a closure
	var isValid func(int, int) bool
	isValid = func(row, col int) bool {
		// 检查列是否有冲突
		for i := range cb {
			if cb[i][col] == 'Q' {
				return false
			}
		}

		n := len(cb)
		// 检查右上是否有皇后冲突
		// note:
		// i = row -1, j = col + 1; i >= 0 && j < 0; i--, j++
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {

			if cb[i][j] == 'Q' {
				return false
			}
		}

		// 检查坐上角是否有皇后冲突
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if cb[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	// backtrack function, a closure
	var btf func(int)
	btf = func(row int) {
		// 触发结束条件
		if row == len(cb) {
			res = append(res, cb.Copy())
			return
		}

		// 列举所有可能的决策
		for col := range cb[row] {
			// 排除不合法的选择
			if !isValid(row, col) {
				continue
			}

			// 做选择
			cb[row][col] = 'Q'

			// 进入下一行决策
			btf(row + 1)

			// 撤销选择
			cb[row][col] = '.'

		}

	}

	btf(0) // 从0开始，回溯每一层的决策

	return res
}
