package main

import "fmt"

func main() {
	title, pages := getBookInfo()
	fmt.Println(title)	
	fmt.Println(pages)

	red := color("red")
	fmt.Println(red.describe("is an awesome color"))

/* 	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, cardValue := range cardValues {
		for _, cardSuit := range cardSuits {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	cards.display() */
}


