package main

import (
	"fmt"
)

//j []int，j[n]存储从0跳到n的最少跳数, j[0]=0
//那么j[n+1]:
//	i:0->n，所有能一步从i跳到n+1的j[i]中的最小值加1, 如果n+1不可达则为-1
func jump1(nums []int) int {
	count := make([]int, len(nums))
	count[0] = 0
	for i := 1; i < len(count); i++ {
		var min int = -1
		for j := 0; j < i; j++ {
			if nums[j] >= i-j && count[j] != -1 {
				if min == -1 || min > count[j] {
					min = count[j]
				}
			}
		}
		if min == -1 {
			count[i] = -1
		} else {
			count[i] = min + 1
		}
	}
	return count[len(count)-1]
}

// count []int, count[n]记录从0跳到n最少的步数，count[0]=0
// nums[0]的值为m,且m>0 则从0可以一步跳到0+1,0+2,...0+m任意一个位置,  那么count[0+1], count[0+2]...,count[0+m]的值为j[0]+1=1
// nums[1]的值为n,且n>0 则从1可以一步跳到1+1,1+2,...1+n， 如果对应位置有值，说明之前用更少的步骤就能跳到该位置，则忽略
// 否则就设置为j[1]+1
// 以此类推
// 优化，当第一次跳到了len(nums)-1，则可定是最少步骤了，直接返回
func jump(nums []int) int {
	count := make([]int, len(nums))
	count[0] = 0
	for i := 0; i < len(nums); i++ {
		if n := nums[i]; n > 0 {
			for j := i + 1; j <= i+n && j < len(nums); j++ {
				if count[j] == 0 {
					count[j] = count[i] + 1
					if j == len(nums)-1 {
						return count[j]
					}
				}
			}
		}
	}
	return count[len(nums)-1]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2}))
}
