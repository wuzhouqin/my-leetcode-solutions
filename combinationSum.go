package main

import (
	// "fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	return combinationSumBottomUp(candidates, target)
}

func combinationSumBottomUp(candidates []int, target int) [][]int {
	var solution map[int][][]int = map[int][][]int{0: {{}}}
	for n := 1; n <= target; n++ {
		var sn [][]int = nil
		for _, c := range candidates {
			if n-c < 0 {
				continue
			}
			if base, ok := solution[n-c]; ok {
				for _, sub := range base {
					var s []int = make([]int, len(sub)+1)
					copy(s, sub)
					s[len(s)-1] = c
					sort.Slice(s, func(i, j int) bool {
						return s[i] < s[j]
					})
					if !contains(sn, s) {
						sn = append(sn, s)
					}
				}
			}
		}
		if sn != nil {
			solution[n] = sn
			//fmt.Println(n, solution[n])
		}
	}

	return solution[target]
}

func combinationSumTopDown(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	r, _ := combination(candidates, target)
	u := uniq(r)
	// for _, sub := range u {
	// 	fmt.Println(sub)
	// }
	return u
}

func uniq(s [][]int) [][]int {
	u := make([][]int, 0)
	for _, check := range s {
		sort.Slice(check, func(i, j int) bool {
			return check[i] < check[j]
		})
		if !contains(u, check) {
			u = append(u, check)
		}
	}
	return u
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func contains(set [][]int, target []int) bool {
	for _, e := range set {
		if equal(e, target) {
			return true
		}
	}

	return false
}

func combination(candidates []int, target int) ([][]int, bool) {
	if target < candidates[0] {
		return nil, false
	}

	solution := make([][]int, 0)
	for i := 0; i < len(candidates); i++ {
		left := target - candidates[i]
		if left < 0 {
			continue
		}
		if left == 0 {
			solution = append(solution, []int{candidates[i]})
		} else {
			s, ok := combination(candidates, left)
			if ok {
				for j := range s {
					s[j] = append(s[j], candidates[i])
				}
				solution = append(solution, s...)
			}
		}
	}

	return solution, len(solution) > 0
}

// func main() {
// 	combinationSum([]int{2, 3, 5}, 8)
// }
