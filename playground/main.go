package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Stores number and its index

	for i, num := range nums {
		complement := target - num
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return []int{} // No solution
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
}