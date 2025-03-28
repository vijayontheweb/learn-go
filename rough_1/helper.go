package main

import (
	"fmt"
	"strconv"
)

type deck []string

func (d deck) display() {
	for index, card := range d {
		fmt.Println(strconv.Itoa(index+1)+". "+card)
	}
}