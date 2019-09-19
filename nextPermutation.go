package main

import (
	// "fmt"
	"sort"
)

func nextPermutation(nums []int) {
	var lastAscIndex int = -1
	var toSwitchIndex = -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			lastAscIndex = i
			toSwitchIndex = i + 1
		} else {
			if toSwitchIndex != -1 && nums[i+1] < nums[toSwitchIndex] && nums[i+1] > nums[lastAscIndex] {
				toSwitchIndex = i + 1
			}
		}
	}
	if lastAscIndex != -1 {
		nums[lastAscIndex], nums[toSwitchIndex] = nums[toSwitchIndex], nums[lastAscIndex]
	}
	toSort := nums[lastAscIndex+1:]
	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i] < toSort[j]
	})
}

// func main() {
// 	nums := []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
// 	nextPermutation(nums)
// 	fmt.Println(nums)
// }
