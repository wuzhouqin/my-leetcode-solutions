package main

import (
	// "fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := "1"
	if n == 1 {
		return s
	}
	for i := 2; i <= n; i++ {
		read := make([]byte, 0)
		char := s[0]
		count := 1
		for j := 1; j < len(s); j++ {
			if char != s[j] {
				read = append(read, strconv.Itoa(count)...)
				read = append(read, char)
				char = s[j]
				count = 1
			} else {
				count++
			}
		}
		read = append(read, strconv.Itoa(count)...)
		read = append(read, char)
		s = string(read)
	}

	return s
}

// func main() {
// 	for i := 1; i < 31; i++ {
// 		fmt.Println(countAndSay(i))
// 	}
// }
