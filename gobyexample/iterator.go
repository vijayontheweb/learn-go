package main

import "fmt"

// List is a generic type that holds a slice of elements.
type List[T any] struct {
    elements []T
}

// Push adds an element to the list.
func (l *List[T]) Push(value T) {
    l.elements = append(l.elements, value)
}

// All returns a channel to iterate over all elements in the list.
func (l *List[T]) All() <-chan T {
    ch := make(chan T)
    go func() {
        for _, e := range l.elements {
            ch <- e
        }
        close(ch)
    }()
    return ch
}

func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    for e := range lst.All() {
        fmt.Println(e)
    }
}