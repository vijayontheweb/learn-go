package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, cardValue := range cardValues {
		for _, cardSuit := range cardSuits {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

func (d deck) print(description ...string) {
	if len(description) > 0 {
		fmt.Println(description[0])
		fmt.Println(strings.Repeat("-", len(description[0])))
	} else {
		fmt.Println(strings.Repeat("-", 100))
	}
	for index, card := range d {
		fmt.Println(strconv.Itoa(index+1) + ". " + card)
	}
}

func deal(cards deck, number int) (deck, deck) {
	return cards[:number], cards[number:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) string{
	var deckStr string
	fileInBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error While Reading File")
		//os.Exit(0);
	} else {
		deckStr = string(fileInBytes)
		fmt.Println(deckStr)		
	}
	return deckStr
}

func stringToDeck(deckString string) deck {
	deckStringArr := strings.Split(deckString, ",")
	return deck(deckStringArr)	
}

func (cards deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	size := len(cards)
	for currIndex := range cards{
		newIndex := random.Intn(size)
		cards[currIndex],cards[newIndex] = cards[newIndex],cards[currIndex]
	}
	/*
	for currIndex := 0; currIndex < size; currIndex++ {
		newIndex := random.Intn(size)
		temp := cards[currIndex]
		cards[currIndex] = cards[newIndex]
		cards[newIndex] = temp
    }
	*/	
}
