package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	sum := 0
	CalculateBranchSum(root, &sum)
	return sum
}
func CalculateBranchSum(item *BinaryTree, sum *int) {
	fmt.Println("sum is {}", *sum)
	if item.Left != nil {
		*sum += 1
		CalculateBranchSum(item.Left, sum)
	}
	if item.Right != nil {
		*sum += 1
		CalculateBranchSum(item.Right, sum)
	}
}
func main() {
	bit_tree := BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}

	fmt.Println("{}", NodeDepths(&bit_tree))
}
