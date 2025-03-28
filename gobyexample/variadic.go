package main

import "fmt"

func sum(nums ...int) int {
	//use lambda function to calculate sum of all numbers
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
	fmt.Println(sum(1, 2, 3)) // Output: 6
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...)) // Output: 10
}


