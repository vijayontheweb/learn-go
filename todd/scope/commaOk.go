package main

import "fmt"

func main() {
	someMap := map[string]int{"one": 1, "two": 2}
	first, ok := someMap["one"]
	if ok {
		fmt.Println("Value of key 'one' is", first)
	} else {
		fmt.Println("Key 'one' not present in the map")
	}
	third, ok := someMap["three"]
	if ok {
		fmt.Println("Value of key 'three' is", third)
	} else {
		fmt.Println("Key 'three' not present in the map")
	}
	if fourth, ok := someMap["four"]; ok {
		fmt.Println("Value of key 'four' is", fourth)
	} else {
		fmt.Println("Key 'four' not present in the map")
	}


	ch := make(chan int, 2)
	ch <- 10
	close(ch)
  
	value, ok := <-ch
	fmt.Println("Value:", value, "Channel Open:", ok)
  
	value, ok = <-ch
	fmt.Println("Value:", value, "Channel Open:", ok)
}
