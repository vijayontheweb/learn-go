package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate a random number and assign to num
	num := rand.Intn(5)
	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Unknown Number")
	}

	//use commas to separate multiple expressions in the same case statement
	//use the default case to catch all other cases
	//use the break statement to exit the switch statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way to express if-else logic
	//Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	//use type switch to determine the type of an interface value
	//In this example, the variable t will have the type of the interface value in the case statement
	//The case statement will be executed if the type of the interface value is time.Time
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		case time.Time:
			fmt.Println("I'm a time.Time")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(time.Now())
	whatAmI(3.14)
	
}
