package main

import (
	"fmt"
	str "strings"
)

var p = fmt.Println

func romanToInt(s string) int {
	info := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var output int
	Arr := str.Split(s, "")

	//output = output + info[Arr[len(Arr)-1]]

	for i := len(Arr) - 1; i > 0; i-- {

		if info[Arr[i-1]] >= info[Arr[i]] {
			output = output + info[Arr[i]]

		} else {
			output = output - info[Arr[i]]

		}
		p(output)

	}
	output = output + info[Arr[0]]
	return output

}

func main() {
	var roman = "MCMXCIV"

	p(romanToInt(roman))

}

// not correctly solved
