package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s := "hello"
	fmt.Println(utf8.RuneCountInString(s))
}