package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var output []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i && nums[i]+nums[j] == target {
				output = append(output, i, j)
				return output
			}
		}
	}
	return output
}

func main() {

	num := []int{-1, -2, -3, -4, -5}
	fmt.Println(twoSum(num, -8))
}
