package main

import "fmt"

func main() {
	var empty map[string]string
	fmt.Println(empty)

	//hexcodes
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for color", color, "is", hex)
	}
}
