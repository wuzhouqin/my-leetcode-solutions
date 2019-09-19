package main

// import (
// 	"fmt"
// )

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return searchRangeIn(nums, 0, len(nums)-1, target)
}

func searchRangeIn(nums []int, low int, high int, target int) []int {
	if high == low {
		if nums[low] == target {
			return []int{low, low}
		} else {
			return []int{-1, -1}
		}
	}

	mid := (high + low) / 2
	if nums[mid] < target {
		return searchRangeIn(nums, mid+1, high, target)
	} else if nums[mid] == target {
		min, max := mid-1, mid+1
		for min >= 0 {
			if nums[min] == target {
				min--
			} else {
				break
			}
		}
		for max < len(nums) {
			if nums[max] == target {
				max++
			} else {
				break
			}
		}
		return []int{min + 1, max - 1}
	} else {
		return searchRangeIn(nums, low, mid, target)
	}
}

// func main() {
// 	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
// }
