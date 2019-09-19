package main

import (
	"fmt"
)

func divide(dividend int, divisor int) int {
	const MIN int32 = -1 << 31
	const MAX int32 = ^MIN
	a, b := int32(dividend), int32(divisor)

	negtive := (a < 0) != (b < 0)
	if b == MIN {
		if a == MIN {
			return 1
		} else {
			return 0
		}
	}

	var result int = 0
	if b < 0 {
		b = -b
	}

	if a == MIN {
		if b == 1 {
			if negtive {
				return int(MIN)
			} else {
				return int(MAX)
			}
		}
		a = -b - a
		result += 1
	} else if a < 0 {
		a = -a
	}
	fmt.Println(a, b)
	for a >= b {
		for i := uint(0); ; i++ {
			temp := b << i
			if a-temp < temp {
				result += 1 << i
				a = a - temp
				break
			}
		}
	}

	if negtive {
		result = -result
	}

	return result
}

// func main() {
// 	r := divide(-2147483648, 2)
// 	fmt.Println(r)
// }
