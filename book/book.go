package main
 
import "fmt"
 
type book string
 
func (b book) printTitle() {
    fmt.Println(b)
}
 
func main() {
    var b book = "Harry Potter"
    b.printTitle()
    arr := [5] int {5,4,3,2,1}
    for index, number := range arr {
        fmt.Println(index, number)
    }
    
}