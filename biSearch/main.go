package main

import (
	"errors"
	"fmt"
)

func BinarySearch(arr []int, l int, h int, key int) (int, error) {

	if l <= h {
		mid := l + (h-l)/2

		if arr[mid] == key {
			return mid, nil
		}
		if arr[mid] > key {
			return BinarySearch(arr, l, mid-1, key)
		}

		return BinarySearch(arr, mid+1, h, key)

	} else {
		return 0, errors.New("error")
	}
}

func main() {

	numbers := [10]int{4, 5, 6, 8, 13, 15, 22, 56, 78, 89}

	result, err := BinarySearch(numbers[:], 0, len(numbers)-1, 778)
	if err != nil {
		fmt.Print("Not Found")
		return
	}
	fmt.Printf("Found at %v", result)

}
