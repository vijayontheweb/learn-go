package main

import "fmt"

type contactInfo struct {
	addressLine1 string
	addressLine2 string
	city         string
	zipCode      int
}

type personWithAddress struct {
	firstName string
	lastName  string
	address   contactInfo
}

type personWithMinimalAddress struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	type person struct {
		firstName string
		lastName  string
	}

	instance1 := person{firstName: "Vijay", lastName: "Anand"}
	fmt.Println(instance1)

	var instance2 person
	fmt.Println(instance2)
	fmt.Printf("%+v\n", instance2)
	var instance3 person
	instance3.firstName = "Priyanka"
	instance3.lastName = "KV"
	fmt.Println(instance3)
	fmt.Printf("%+v\n", instance3)

	embeddedInstance := personWithAddress{
		firstName: "Gautham",
		lastName:  "VA",
		address: contactInfo{
			addressLine1: "102, Jacksonville",
			addressLine2: "Harvard Street",
			city:         "Nevada",
			zipCode:      12312,
		},
	}
	fmt.Println(embeddedInstance)

	embeddedInstance2 := personWithMinimalAddress{
		firstName: "Meenakshi",
		lastName:  "VA",
		contactInfo: contactInfo{
			addressLine1: "102, Jacksonville",
			addressLine2: "Harvard Street",
			city:         "Nevada",
			zipCode:      12312,
		},
	}
	fmt.Println(embeddedInstance2)
	embeddedInstance2.printAll()
	embeddedInstance2.updateCityNonPointer()
	embeddedInstance2.printAll()
	embeddedInstance2Pointer := &embeddedInstance2
	fmt.Printf("pointer -> %p\n", embeddedInstance2Pointer)
	embeddedInstance2Pointer.updateCityPointer()
	embeddedInstance2.printAll()

	//Slice
	sentence := []string{"Hi", "How", "Are", "You"}
	sentence[0] = "Aye"
	fmt.Println("slice -> ", sentence)

	name := "Bill"
	fmt.Println(*&name)

	name = "Gates"
	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p personWithMinimalAddress) printAll() {
	fmt.Println("receiver -> ", p)
}

func (p personWithMinimalAddress) updateCityNonPointer() {
	p.contactInfo.city = "Texas"
}

func (p *personWithMinimalAddress) updateCityPointer() {
	(*p).contactInfo.city = "Buffalo"
}
