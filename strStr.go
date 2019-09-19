package main

// import "fmt"

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	result := -1
	switch {
	case m == 0:
		result = 0
	case m > n:
		result = -1
	default:
		for i := 0; i < n-m+1; i++ {
			found := true
			for j := 0; j < m; j++ {
				if haystack[i+j] != needle[j] {
					found = false
					break
				}
			}
			if found {
				result = i
				break
			}
		}
	}

	return result
}

// func main() {
// 	fmt.Println(strStr("hello", "ll"))
// 	fmt.Println(strStr("aaaaaa", "bba"))
// 	fmt.Println(strStr("asdfas中文asdfa", "中"))
// }
