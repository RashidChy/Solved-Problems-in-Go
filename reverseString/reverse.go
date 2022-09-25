package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {

	Arr := strings.Split(str, "")
	var n int = len(Arr) / 2
	end := len(Arr) - 1

	for i := 0; i < n; i++ {
		temp := Arr[end]
		Arr[end] = Arr[i]
		Arr[i] = temp
		end--
	}
	str = strings.Join(Arr, "")
	return str

}

func main() {

	fmt.Println(Reverse("sakin"))
}
