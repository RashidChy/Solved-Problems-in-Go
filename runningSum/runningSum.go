package main

import (
	"errors"
	"fmt"
)

func runningSum(nums []int) ([]int, error) {
	length := len(nums)
	result := []int{}
	var count int
	var sum int
	if nums != nil {
		for i := 0; i < length; i++ {
			for i := 0; i <= count; i++ {

				if count != length {
					sum = sum + nums[i]

				}
			}
			result = append(result, sum)
			sum = 0
			count++

		}
		return result, errors.New("no error")
	} else {
		return result, errors.New("error")
	}

}

func main() {

	num := [4]int{1, 2, 3, 4}
	fmt.Println(runningSum(num[:]))
}
