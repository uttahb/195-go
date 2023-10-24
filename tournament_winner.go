package main

import (
	"fmt"
)

func main() {
	array := [][]string{{"html", "c#"},
		{"c#", "python"},
		{"python", "html"}}
	results := []int{0, 0, 1}
	fmt.Println(TournamentWinner(array, results))

}
func TournamentWinner(array [][]string, results []int) string {
	// Write your code here.
	dict := map[string]int{}
	max := ""
	max_count := 0
	for i, v := range array {
		res := 1
		if results[i] == 1 {
			res = 0
		}
		label := v[res]
		_, ok := dict[label]
		if ok {
			dict[label] += 1
		} else {
			dict[label] = 1
		}
		if max_count < dict[label] {
			max_count = dict[label]
			max = label
		}
	}
	return max
}
