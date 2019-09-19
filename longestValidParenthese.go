package main

// import (
// 	"fmt"
// )

type item struct {
	character byte
	index     int
}

func longestValidParentheses(s string) int {
	var n int = len(s)
	if n == 0 {
		return 0
	}

	max := 0
	stack := make([]item, 0)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, item{'(', i})
		} else {
			currLen := len(stack)
			if currLen > 0 {
				peek := stack[currLen-1]
				if peek.character == '(' {
					if currLen > 1 {
						max = maxOf(max, i-stack[currLen-2].index)
					} else {
						max = maxOf(max, i+1)
					}
					stack = stack[:currLen-1]
				} else {
					stack = append(stack, item{')', i})
				}
			} else {
				stack = append(stack, item{')', i})
			}

		}
	}

	return max
}

func minOf(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func maxOf(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// func main() {
// 	fmt.Println(longestValidParentheses("(()"))
// 	fmt.Println(longestValidParentheses(")()())"))
// }
