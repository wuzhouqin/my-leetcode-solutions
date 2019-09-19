package main

// import (
// 	"fmt"
// )

func findSubstring(s string, words []string) []int {
	var result []int = nil
	if len(words) == 0 {
		return result
	}
	wordLen := len(words[0])
	subLen := len(words) * wordLen
	strLen := len(s)
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word] = wordMap[word] + 1
	}

	for i := 0; i < strLen-subLen+1; i++ {
		seenMap := make(map[string]int)
		for j := 0; j < len(words); j++ {
			start := i + j*wordLen
			seenWord := s[start : start+wordLen]
			n, ok := wordMap[seenWord]
			if !ok {
				break
			}
			seenMap[seenWord] = seenMap[seenWord] + 1
			if n < seenMap[seenWord] {
				break
			}
			if j == len(words)-1 {
				result = append(result, i)
			}
		}
	}
	return result
}

// func main() {
// 	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
// 	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
// }
