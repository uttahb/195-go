package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(TwoNumberSum(array, 7))
}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	sort.Ints(array)
	lp := 0
	rp := len(array) - 1
	for {
		sum := array[lp] + array[rp]
		if sum < target {
			lp += 1
		} else if sum > target {
			rp -= 1
		} else {
			return []int{array[lp], array[rp]}
		}
		if lp > rp {
			break
		}
	}

	return []int{}
}
