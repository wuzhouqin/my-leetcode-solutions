package main

import (
	// "fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	seen := make(map[string]int)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			rowKey := "row" + strconv.Itoa(i) + string(board[i][j])
			_, ok := seen[rowKey]
			if ok {
				return false
			}
			seen[rowKey] = 1

			columnKey := "column" + strconv.Itoa(j) + string(board[i][j])
			_, ok = seen[columnKey]
			if ok {
				return false
			}
			seen[columnKey] = 1

			cubeKey := "cube" + strconv.Itoa(i/3+1) + strconv.Itoa(j/3+1) + string(board[i][j])
			_, ok = seen[cubeKey]
			if ok {
				return false
			}
			seen[cubeKey] = 1
		}
	}

	return true
}

// func main() {
// 	nums := [][]byte{
// 		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	}
// 	fmt.Println(isValidSudoku(nums))
// }
