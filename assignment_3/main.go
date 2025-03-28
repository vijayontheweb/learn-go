package main

import (
	"fmt"
	"io"
	"os"
)

type customWriter struct {
	
}

func (w customWriter) Write(p []byte) (n int, err error){
	fmt.Println(string(p))
	return len(p),nil
}

func main() {
	// Get command-line arguments
	args := os.Args

	// Print all arguments
	fmt.Println("Command-line arguments:")
	for i, arg := range args {
		fmt.Printf("Arg %d: %s\n", i, arg)
	}
	//fileContents := readFromFile(os.Args[1])
	//fmt.Println(fileContents)
	f,err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error While Reading File")
		os.Exit(0);
	} else {
		writer := customWriter{}
		io.Copy(writer,f)
	}

}

func readFromFile(filename string) string{
	var deckStr string
	fileInBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error While Reading File")
		os.Exit(0);
	} else {
		deckStr = string(fileInBytes)		
	}
	return deckStr
}
