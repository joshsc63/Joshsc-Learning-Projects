package main

import "fmt"

func main () {
	
	//var colors map[string]string
	//colors := make(map[string]string)
	//colors["white"] = "#fffff"
	
	//declaring map where keys are string & values are string
	colors := map[string]string{
		"red":"#ff000",
		"green":"#4bf745",
		"white":"#fffff",
	}

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}