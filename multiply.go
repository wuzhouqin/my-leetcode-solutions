package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	result := make([]byte, 220)
	clear(result)
	temp := make([]byte, 220)
	for i := len(num2) - 1; i >= 0; i-- {
		if num2[i] == '0' {
			continue
		}
		clear(temp)
		multiplyOne(num1, num2[i], len(temp)-(len(num2)-i), temp)
		add(result, temp)
	}

	var i int = 0
	for {
		if i == len(result) {
			return "0"
		}
		if result[i] != '0' {
			break
		} else {
			i++
		}
	}
	return string(result[i:])
}

func clear(a []byte) {
	for i := 0; i < len(a); i++ {
		a[i] = '0'
	}
}

// a, b are same length, and long enough to hold the result
// store result to a then b is free to be recycle
// a, b are left padding with '0'
func add(a, b []byte) {
	var carrior byte = 0
	for i := len(a) - 1; i >= 0; i-- {
		c := (a[i] - '0') + (b[i] - '0') + carrior
		a[i] = (c % 10) + '0'
		carrior = c / 10
	}
}

func multiplyOne(num1 string, b byte, offset int, holder []byte) {
	var carrior byte = 0
	for i := len(num1) - 1; i >= 0; i-- {
		c := (num1[i]-'0')*(b-'0') + carrior
		holder[offset] = (c % 10) + '0'
		carrior = c / 10
		offset--
	}
	if carrior > 0 {
		holder[offset] = carrior + '0'
	}
}

func main() {
	fmt.Println(multiply("999", "999"))
	// fmt.Println(multiply("0", "33"))
	// fmt.Println(multiply("2", "3"))
	// fmt.Println(multiply("123", "456"))
}
