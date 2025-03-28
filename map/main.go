package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":  "R",
		"blue": "B",
	}
	fmt.Println(colors)
	var emptyMap1 map[string]string
	fmt.Println(emptyMap1)
	emptyMap2 := make(map[string]string)
	fmt.Println(emptyMap2)
	newMap := make(map[string]string)
	newMap["animal"] = "tiger"
	newMap["bird"] = "eagle"
	newMap["insect"] = "ant"
	newMap["reptile"] = "crocodile"
	fmt.Println(newMap)
	delete(newMap, "bird")
	fmt.Println(newMap)
	printMap(newMap)
}

func printMap(mapInstance map[string]string) {
	for key, value := range mapInstance {
		fmt.Printf("key->%v value->%v\n", key, value)
	}
}
