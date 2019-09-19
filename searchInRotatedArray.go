package main

// import (
// 	"fmt"
// )

func search(nums []int, target int) int {
	return searchIn(nums, 0, len(nums)-1, target)
}

func searchIn(nums []int, start int, end int, target int) int {
	// fmt.Println("search in ", start, end)
	if (end - start + 1) < 4 {
		for i := start; i <= end; i++ {
			if target == nums[i] {
				return i
			}
		}
		return -1
	}
	// split
	x1, x2 := start, (end+start)/2
	x3, x4 := x2+1, end
	if nums[x1] < nums[x2] {
		if target >= nums[x1] && target <= nums[x2] {
			return searchIn(nums, x1, x2, target)
		} else {
			return searchIn(nums, x3, x4, target)
		}
	} else {
		// nums[x3:x4+1] must be asc order
		if target >= nums[x3] && target <= nums[x4] {
			return searchIn(nums, x3, x4, target)
		} else {
			return searchIn(nums, x1, x2, target)
		}
	}
}

// func main() {
// 	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
// }
