package main

import "testing"

func TestBiSearch(t *testing.T) {
	numbers := [10]int{4, 5, 6, 8, 13, 15, 22, 56, 78, 89}
	got, err := BinarySearch(numbers[:], 0, len(numbers)-1, 13)
	if err == nil {
		t.Errorf("test passed failed")
	}
	got, err = BinarySearch(numbers[:], 0, len(numbers)-1, 4)
	if err != nil {
		t.Errorf("test failed")
	} else {
		if got != 0 {
			t.Errorf("Index Wrong, expected: 0, found: %v", got)
		}
	}
}
