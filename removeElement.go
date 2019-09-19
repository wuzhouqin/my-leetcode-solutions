package main

// import "fmt"

func removeElement(nums []int, val int) int {
	var eleToCheckIndex int = 0
	var sameValueIndex int = len(nums)

	for eleToCheckIndex < sameValueIndex {
		if nums[eleToCheckIndex] == val {
			nums[eleToCheckIndex], nums[sameValueIndex-1] = nums[sameValueIndex-1], nums[eleToCheckIndex]
			sameValueIndex--
		} else {
			eleToCheckIndex++
		}
	}
	return sameValueIndex
}

// func main() {
// 	nums := []int{3, 2, 2, 3}
// 	n := removeElement(nums, 3)
// 	fmt.Println(nums, n)

// 	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
// 	n = removeElement(nums, 2)
// 	fmt.Println(nums, n)
// }
