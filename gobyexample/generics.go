package main

import "fmt"

func main(){
	fmt.Println(max(1, 2)) // Output: 2
	fmt.Println(max(1.1, 2.2)) // Output: 2.2
}

func max[T int|float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}