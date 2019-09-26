package main

// "fmt"

func isMatch(s string, p string) bool {
	// 有连续的*的时候，只保留一个
	t := []byte{}
	for i := 0; i < len(p); i++ {
		if p[i] == '*' && len(t) > 0 && t[len(t)-1] == '*' {
			continue
		} else {
			t = append(t, p[i])
		}
	}
	// 空p只匹配空字符串
	if len(t) == 0 {
		return len(s) == 0
	}

	p = string(t)
	set := map[int]bool{0: true}
	for i := 0; i < len(s); i++ {
		newSet := make(map[int]bool)
		for j, _ := range set {
			if j >= len(p) {
				continue
			}
			switch p[j] {
			case '*':
				newSet[j] = true
				newSet[j+1] = true
				if j+1 < len(p) {
					if p[j+1] == '?' || p[j+1] == s[i] {
						newSet[j+2] = true
					}
				}
			case '?':
				newSet[j+1] = true
			default:
				if p[j] == s[i] {
					newSet[j+1] = true
				}
			}
		}
		if len(newSet) == 0 {
			return false
		} else {
			set = newSet
		}
	}

	for j, _ := range set {
		if j == len(p) {
			return true
		}
		if j == len(p)-1 && p[j] == '*' {
			return true
		}
	}

	return false
}

// func main() {
// 	fmt.Println(isMatch("ho", "**ho*"))
// }
