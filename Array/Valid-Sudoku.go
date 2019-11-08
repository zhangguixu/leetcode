package main

import (
	"fmt"
)

func main() {
	data := make([][]byte, 9)

	data[0] = []byte("53..7....")
	data[1] = []byte("6..195...")
	data[2] = []byte(".98....6.")
	data[3] = []byte("8...6...3")
	data[4] = []byte("4..8.3..1")
	data[5] = []byte("7...2...6")
	data[6] = []byte(".6....28.")
	data[7] = []byte("...419..5")
	data[8] = []byte("....8..79")

	// data[0] = []byte("......5..")
	// data[1] = []byte(".........")
	// data[2] = []byte(".........")
	// data[3] = []byte("93..2.4..")
	// data[4] = []byte("..7...3..")
	// data[5] = []byte(".........")
	// data[6] = []byte("...34....")
	// data[7] = []byte(".....3...")
	// data[8] = []byte(".....52..")
	fmt.Println(isValidSudoku(data))
	fmt.Println(isValidSudoku_v1(data))
}

/*
	数独的校验
	满足三个规则即可
	Each row must contain the digits 1-9 without repetition.
	Each column must contain the digits 1-9 without repetition.
	Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

	题目中给定的尺寸都是 9 * 9，因此这里就无需校验的

	保证1-9不重复
	[
		["5","3",".",".","7",".",".",".","."],
		["6",".",".","1","9","5",".",".","."],
		[".","9","8",".",".",".",".","6","."],
		["8",".",".",".","6",".",".",".","3"],
		["4",".",".","8",".","3",".",".","1"],
		["7",".",".",".","2",".",".",".","6"],
		[".","6",".",".",".",".","2","8","."],
		[".",".",".","4","1","9",".",".","5"],
		[".",".",".",".","8",".",".","7","9"]
	]
*/

/*
	遍历，这种算法需要遍历 9 * 9，并且需要额外的空间存储相同元素出现的位置信息
*/

type Position struct {
	X int
	Y int
}

func isRepeated(p1 *Position, p2 *Position) bool {
	if p1.X == p2.X || p1.Y == p2.Y || (p1.X/3 == p2.X/3 && p1.Y/3 == p2.Y/3) {
		return true
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	numMap := make(map[byte][]*Position, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			b := board[i][j]
			list, ok := numMap[b]
			pos := &Position{i, j}
			if !ok {
				list = make([]*Position, 0)
			} else {
				for _, p := range list {
					if isRepeated(pos, p) {
						return false
					}
				}
			}
			list = append(list, pos)
			numMap[b] = list
		}
	}
	return true
}

/*
	disucss most votes的第一名，思路是一致的，只是使用自己的编码来存储，会更加节省空间
	Collect the set of things we see, encoded as strings. For example:

	'4' in row 7 is encoded as "(4)7".
	'4' in column 7 is encoded as "7(4)".
	'4' in the top-right block is encoded as "0(4)2"

	额，这种实现的runtime好长，执行效率比较低
*/
func isValidSudoku_v1(board [][]byte) bool {
	m := make(map[string]int, 27)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			rStr := string([]byte{'(', b, ')', byte(i + 48)})
			cStr := string([]byte{byte(j + 48), '(', b, ')'})
			bStr := string([]byte{byte(i/3 + 48), '(', b, ')', byte(j/3 + 48)})
			if m[rStr] == 1 || m[cStr] == 1 || m[bStr] == 1 {
				return false
			}
			m[rStr]++
			m[cStr]++
			m[bStr]++
		}
	}
	return true
}

/*
	运用位操作来记录状态
*/
func isValidSudoku_v2(board [][]byte) bool {
	rowBit, colBit := 0, make([]int, len(board))
	for i := 0; i < len(board); i++ {
		rowBit = 0
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				bit := 1 << uint(board[i][j]-'0'-1)
				if rowBit&bit > 0 || colBit[j]&bit > 0 {
					return false
				}
				if !isValidCube(board, i, j) {
					return false
				}
				rowBit |= bit
				colBit[j] |= bit
			}
		}
	}

	return true
}

func isValidCube(board [][]byte, row, col int) bool {
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if i == row && j == col {
				continue
			}
			if board[i][j] == board[row][col] {
				return false
			}
		}
	}
	return true
}
