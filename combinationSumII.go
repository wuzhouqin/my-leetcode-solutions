package main

import (
	// "fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	return combinationSum2BottomUp(candidates, target)
}

func combinationSum2TopDown(candidates []int, target int) ([][]int, bool) {
	if target < 0 {
		return nil, false
	}
	if target == 0 {
		return [][]int{{}}, true
	}
	if len(candidates) == 0 {
		return nil, false
	}

	value := candidates[0]
	var solutions [][]int = nil
	if sub, ok := combinationSum2TopDown(candidates[1:], target); ok {
		for _, s := range sub {
			solutions = append(solutions, s)
		}
	}

	if sub, ok := combinationSum2TopDown(candidates[1:], target-value); ok {
		for _, s := range sub {
			s = append(s, value)
			sort.Slice(s, func(i, j int) bool {
				return s[i] < s[j]
			})
			if !contains(solutions, s) {
				solutions = append(solutions, s)
			}
		}
	}

	return solutions, solutions != nil
}

func combinationSum2BottomUp(candidates []int, target int) [][]int {
	solutions := map[int][][]int{0: {{}}}
	for i := 0; i < len(candidates); i++ {
		// suppose the candidates are cadidates[:i+1]
		// all possibilities are
		//     1 pick cadidates[i]
		// 	   2 not pick cadidates[i]
		// base on solutions of when candidates are candidates[:i]
		add := make(map[int][][]int)
		for key := range solutions {
			t := key + candidates[i]
			if t > target {
				continue
			}
			for _, s := range solutions[key] {
				ns := make([]int, len(s)+1)
				copy(ns, s)
				ns[len(ns)-1] = candidates[i]
				add[t] = append(add[t], ns)
			}
		}

		for key := range add {
			for _, ns := range add[key] {
				sort.Slice(ns, func(i, j int) bool {
					return ns[i] < ns[j]
				})
				if !contains(solutions[key], ns) {
					solutions[key] = append(solutions[key], ns)
				}
			}
		}
	}

	// for key := range solutions {
	// 	fmt.Println(key, solutions[key])
	// }
	return solutions[target]
}

// func main() {
// 	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
// }
