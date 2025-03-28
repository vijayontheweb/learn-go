package main

import "fmt"

var deckSize int
//deckSize := 20

func getNumber() int{
	return 5
}

func getTwo() string {
	return "two"
}


func main(){
	//var name string = "vijay"
	name := "anand"
	fmt.Println("hello "+name)
	name = "vijay"
	fmt.Println("hello "+name)
	
	
	fmt.Println(deckSize)
	deckSize = 53
	fmt.Println(deckSize)
	fmt.Println(getNumber())
	printState()
	//Array and Slice
	cards := [] string { "one", getTwo()}
	cards = append(cards, "three")
	for i, card := range cards{
		fmt.Println(i, card)
	}
	cards2 := deck { "four", "five", "six"}
	cards2.print()

	//final
	cards3 := newDeck()
	cards3.print()
	hand, remainingCards := deal(cards3,5)
	hand.print("HAND:")
	remainingCards.print("REMAINING:")
	fmt.Println(hand.toString())
	hand.writeToFile("my_file.txt")	
	readFromFile("some_file.txt")
	deckStr := readFromFile("my_file.txt")
	handCards := stringToDeck(deckStr)
	handCards.print("HAND FROM FILE:")
	handCards.shuffle()
	handCards.print("SHUFFLED HAND:")
}

