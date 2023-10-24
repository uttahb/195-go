package main

import (
	"fmt"
)

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{22, 25, 6}
	fmt.Println(IsValidSubsequence(array, sequence))

}
func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	for _, v := range array {

		if v == sequence[0] {
			sequence = sequence[1:]
			if len(sequence) == 0 {
				return true
			}
		}
	}
	return false
}
