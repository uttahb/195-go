package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{-2, -1}
	fmt.Println(SortedSquaredArray(array))

}
func SortedSquaredArray(array []int) []int {
	// Write your code here.
	for i := 0; i < len(array); i++ {
		*&array[i] = array[i] * array[i]

	}
	sort.Ints(array)
	return array
}
