package main

// import (
// 	"fmt"
// )

func solveSudoku(board [][]byte) {
	solve(board, 0, 0)
}

func solve(board [][]byte, i, j int) bool {

	for m := i; m < 9; m++ {
		n := j
		if m > i {
			n = 0
		}
		for ; n < 9; n++ {
			if board[m][n] == '.' {
				left := choice(board, m, n)
				if len(left) == 0 {
					return false
				} else {
					for c := range left {
						board[m][n] = c
						if solve(board, m, n) {
							return true
						}
					}
					board[m][n] = '.'
					return false
				}
			}
		}
	}

	return true
}

func choice(board [][]byte, i, j int) map[byte]bool {
	left := map[byte]bool{'1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}
	for k := 0; k < 9; k++ {
		if board[i][k] != '.' {
			delete(left, board[i][k])
		}
		if board[k][j] != '.' {
			delete(left, board[k][j])
		}
	}

	ii := i / 3
	jj := j / 3
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			if board[ii*3+m][jj*3+n] != '.' {
				delete(left, board[ii*3+m][jj*3+n])
			}
		}
	}

	return left
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
// 	nums = [][]byte{
// 		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
// 		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
// 		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
// 		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
// 		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
// 		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
// 		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
// 		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
// 		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
// 	}
// 	solveSudoku(nums)
// 	for _, row := range nums {
// 		fmt.Println(string(row))
// 	}
// 	fmt.Println(isValidSudoku(nums))
// }
