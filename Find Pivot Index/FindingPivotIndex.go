package main

import (
	"fmt"
	"time"
)

func pivotIndex(nums []int) int {
	var left, right int

	for i := 0; i < len(nums); i++ {

		go func() {
			for j := i; j > 0; j-- {
				if i == 0 {
					left = 0
				} else {
					left = left + nums[j-1]
				}
			}
		}()

		go func() {
			for j := i; j < len(nums)-1; j++ {

				if i == len(nums)-1 {
					right = 0
				} else {
					right = right + nums[j+1]
				}
			}
		}()

		//fmt.Println("at pivot = ", i)
		//fmt.Println(left, "=", right)
		time.Sleep(500 * time.Millisecond)

		if left == right {
			return i
		} else {
			left = 0
			right = 0
		}
	}
	return -1

}

func main() {
	number := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(number))
}
