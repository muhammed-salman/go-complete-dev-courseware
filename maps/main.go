package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// colors := make(map[int]string)

	// colors[10] = "#ffffff"

	// delete(colors, 10)

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
