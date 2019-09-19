package main

func generateParenthesis(n int) []string {
	holder := make([]rune, 2*n)
	var result []string
	growParenthesis(&result, &holder, 0, 0, 0, n)
	return result
}

func growParenthesis(result *[]string, holderp *[]rune, position int, openCount int, closeCount int, max int) {
	holder := *holderp

	if position == len(holder) {
		*result = append(*result, string(holder))
		return
	}
	if openCount < max {
		holder[position] = '('
		growParenthesis(result, holderp, position+1, openCount+1, closeCount, max)
	}
	if closeCount < openCount {
		holder[position] = ')'
		growParenthesis(result, holderp, position+1, openCount, closeCount+1, max)
	}
}
