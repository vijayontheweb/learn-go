package main

import "fmt"

func main() {
//closure filter by a number
	filter := func(num int) func(int) bool {
		return func(x int) bool {
			return x>num
		}
	}
	//filter by 2
	filterBy2 := filter(2)
	//filter by 3
	filterBy3 := filter(3)
	
	//LIST OF NUMBERS
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(filterNumbers(numbers,filterBy2))
	fmt.Println(filterNumbers(numbers,filterBy3))
}

func filterNumbers(numbers []int, filter func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if filter(num) {
			result = append(result, num)
		}
	}
	return result
}
// Output:
