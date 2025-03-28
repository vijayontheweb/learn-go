package main

func main() {
	//create a slice instead of an array
	var numbers = []int{11, 12, 3:50, 14, 15}
	//append to the slice
	numbers = append(numbers, 16)
	//iterate over the slice
	for index, num := range numbers {
		println(index, num)
	}
	//print the length and capacity of the slice
	println(len(numbers)) // Output: 5
	println(cap(numbers)) // Output: 8
	//create a slice with make
	var slice = make([]int, 5, 10)
	println(len(slice)) // Output: 5
	println(cap(slice)) // Output: 10
	//copying one slice to another
	var slice2 = make([]int, 5)
	copy(slice2, numbers)
	for index, num := range slice2 {	
		println("SLICE2",index, num)
	}
	//slicing a slice
	var slice3 = numbers[1:3]
	for index, num := range slice3 {
		println("SLICE3",index, num)
	}
	
}