package main

import (
	"fmt"
)

func main() {
	data := make([][]byte, 9)

	// data[0] = []byte("53..7....")
	// data[1] = []byte("6..195...")
	// data[2] = []byte(".98....6.")
	// data[3] = []byte("8...6...3")
	// data[4] = []byte("4..8.3..1")
	// data[5] = []byte("7...2...6")
	// data[6] = []byte(".6....28.")
	// data[7] = []byte("...419..5")
	// data[8] = []byte("....8..79")

	data[0] = []byte("......5..")
	data[1] = []byte(".........")
	data[2] = []byte(".........")
	data[3] = []byte("93..2.4..")
	data[4] = []byte("..7...3..")
	data[5] = []byte(".........")
	data[6] = []byte("...34....")
	data[7] = []byte(".....3...")
	data[8] = []byte(".....52..")
	fmt.Println(isValidSudoku(data))
}

/*
	数独的校验
	满足三个规则即可
	Each row must contain the digits 1-9 without repetition.
	Each column must contain the digits 1-9 without repetition.
	Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
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
func isValidSudoku(board [][]byte) bool {
	var numOfrow int
	numOfrow = len(board)
	if numOfrow % 9 != 0 {
		return false
	}
	valid := true
	// 验证行
	for i := 0; i < numOfrow; i++ {
		if !valid {
			break
		}
		row := board[i]
		if len(row) != numOfrow {
			valid = false
			break
		}
		rowMap := make(map[byte]int, 0)
		for j := 0; j < numOfrow; j++ {
			if row[j] == 46 {
				continue
			}
			_, ok := rowMap[row[j]]
			if ok {
				valid = false
				break
			}
			rowMap[row[j]] = 1
		}
	}
	if !valid {
		return false
	}
	// 验证列
	for i := 0; i < numOfrow; i++ {
		if !valid {
			break
		}
		columnMap := make(map[byte]int, 0)
		for j := 0; j < numOfrow; j++ {
			if board[j][i] == 46 {
				continue
			}
			_, ok := columnMap[board[j][i]]
			if ok {
				valid = false
				break
			}
			columnMap[board[j][i]] = 1
		}
	}
	
	if !valid {
		return false
	}

	// 验证3*3方格
	fmt.Println(numOfrow)
	for i := 0; i < numOfrow; i += 3 {
		if !valid {
			break
		}
		for j := 0; j < numOfrow; j += 3 {
			if !valid {
				break
			}
			boardMap := make(map[byte]int, 0)
			for k := 0; k < 3; k++ {
				if !valid {
					break
				}
				for l := 0; l < 3; l++ {
					if board[i + l][j + k] == 46 {
						continue
					}
					_, ok := boardMap[board[i + l][j + k]]
					if ok {
						valid = false
						break
					}
					boardMap[board[i + l][j + k]] = 1
				}
			}
		}
	}

	return valid
}

// height := len(board)
// if height % 9 != 0 {
// 	return false
// } 
// valid := true
// for i := 0; i < height; i++ {
// 	if valid == false {
// 		break
// 	}
// 	row := board[i]
// 	if len(row) != height {
// 		valid = false
// 		break
// 	}
// 	rowMap := make(map[byte]int, 0)
// 	for j := 0; j < height; j++ {
// 		if valid == false {
// 			break
// 		}
// 		if row[j] == 46 {
// 			continue
// 		}
// 		_, ok := rowMap[row[j]]
// 		if ok {
// 			valid = false
// 			break
// 		}
// 		rowMap[row[j]] = 1
// 		columnMap := make(map[byte]int, 0)
// 		for k := 0; k < height; k++ {
// 			if board[k][j] == 46 {
// 				continue
// 			}
// 			_, ok := columnMap[board[k][j]]
// 			if ok {
// 				valid = false
// 				break
// 			}
// 			columnMap[board[k][j]] = 1
// 		}
// 	}
// }
// return valid