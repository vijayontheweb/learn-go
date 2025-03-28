package main

import "fmt"

func main() {
	//If you specify the index with :, the elements in between will be zeroed.
	var numbers = [...]int{11, 12, 3:50, 14, 15}
	fmt.Println(len(numbers)) // Output: 5
	
	for index, num := range numbers {
		fmt.Println(index, num)
	}
	//initializing multi dimensional array
	var multiDimArray = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(multiDimArray)

	

}