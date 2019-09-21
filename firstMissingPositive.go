package main

// "fmt"

// if we have a big enough slice, say flag []bool, all value of flag are init to false
// flag[i] == true means found value i in nums
// then program can scan through nums, and turn on flags

// here is the trick, if we put value v in position nums[v-1],
// and throw away any negitive or zero value or value greater than length of nums,
// the result of comparing nums[v-1] to v can be a flag
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = -1
		}
	}

	for i := 0; i < len(nums); {
		v := nums[i]
		if v == -1 || v == i+1 {
			i++
			continue
		} else {
			// put v to nums[v-1]
			// if nums[v-1] already equals v, v is duplicated, than we simple mark nums[i] to -1
			// else we switch nums[v-1] and nums[i]
			if nums[v-1] == v {
				nums[i] = -1
				i++
			} else {
				nums[v-1], nums[i] = nums[i], nums[v-1]
			}
		}
	}

	// fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

// func main() {
// 	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
// 	fmt.Println(firstMissingPositive([]int{3, 1, 2}))
// 	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
// }
