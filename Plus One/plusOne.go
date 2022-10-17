package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	n:= len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		} 
		digits[i] = 0
	}
	
	digits[0] = 1
	digits = append(digits, 0)

	return digits
}

func main() {
	number := []int{4, 9, 9, 9}
	fmt.Println(plusOne(number))
}
