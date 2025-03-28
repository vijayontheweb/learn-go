package main

import "net/http"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}


	c2 := make(chan string)
	go pass(c2)
	c2 <- "Hi there!"
	

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// ONE WAY
	// for {
	// 	go checkLink(<-c, c)
	// }

	//ALTERNATIVELY
	for l := range c {
	 	go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		c <- link
		return
	}
	c <- link
	println(link, "is up!")
}

func pass(channel chan string) {
	println(<-channel)
}