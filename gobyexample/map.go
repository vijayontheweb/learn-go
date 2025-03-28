package main

import (
	"fmt"
	"maps"
)

func main(){
    //initialize a map
	m := make(map[string]int)
	//set key-value pairs
	m["k1"] = 7
	m["k2"] = 13
	//print the map
	fmt.Println("map:",m) // Output: map: map[k1:7 k2:13]
	fmt.Println("key doesn't exist:",m["k3"])
	fmt.Println("length of map:",len(m))
	//delete a key-value pair
	delete(m,"k2")
	fmt.Println("map after deleting k2:",m)
	value,isPresent := m["k1"]
	fmt.Println("is present k1:",isPresent,"value of k1:",value)
	//equality check
	n := map[string]int{"foo":1,"bar":2}
	nClone := map[string]int{"foo":1,"bar":2}
	fmt.Println("map n:",n)
	fmt.Println("map equality:", maps.Equal(m,n)) 
	fmt.Println("map equality:", maps.Equal(n,nClone)) 
}