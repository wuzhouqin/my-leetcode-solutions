package main

// import (
// 	"fmt"
// )

func searchInsert(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}

	return i
}

// func main() {
// 	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
// 	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
// 	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
// 	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
// }
