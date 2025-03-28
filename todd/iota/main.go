package main

import (
	"fmt"
)

type byteSize int

//enumerated constants are created using the iota enumerator
//KB MB GB TB PB
const (
	_ = iota // 0
	KB byteSize = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) // 1 << (3*10)
	TB = 1 << (iota * 10) // 1 << (4*10)
	PB = 1 << (iota * 10) // 1 << (5*10)
)


func main() {
fmt.Printf("%b %d\n",KB, KB)
fmt.Printf("%b %d\n",MB, MB)
fmt.Printf("%b %d\n",GB, GB)
fmt.Printf("%b %d\n",TB, TB)
fmt.Printf("%b %d\n",PB, PB)
}
