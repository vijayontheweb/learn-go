package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	num := 1 << 2 // shift one to left by 2 bits
	fmt.Printf("%b %d\n",num, num)
}
