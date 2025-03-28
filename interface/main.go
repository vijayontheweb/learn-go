package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type bot interface {
	getGreeting(string, int) (string, error)
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting(name string, id int) (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	if random.Intn(2) == 0 {
		return "Hello " + name + "! Your id is " + strconv.Itoa(id), nil
	} else {
		return "", errors.New("Something unexpected happened :(")
	}

}

func (sb spanishBot) getGreeting(name string, id int) (string, error) {
	return "Hola " + name + "! Tu id es " + strconv.Itoa(id), nil
}

func printGreeting(bot bot, name string) {
	greeting, error := bot.getGreeting(name, 1)
	if error == nil {
		fmt.Println(greeting)
	} else {
		fmt.Println(error)
	}

}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb, "vijay")
	printGreeting(sb, "priya")

}
