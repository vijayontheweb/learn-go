package main

type color string

func getBookInfo() (string, int) {
    return "War and Peace", 1000
}

func (c color) describe(message string) (string){
	return string(c)+" "+message;
}

