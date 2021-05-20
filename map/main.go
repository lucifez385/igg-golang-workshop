package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"set":  33,
		"joe":  290,
		"karn": 30,
	}
	fmt.Println(ages)

	// colors := make(map[string]string)
	// colors["red"] = "#FF0000"
	// colors["black"] = "#00000"

	// delete(colors, "black")
	// fmt.Println(colors)
	// printMap(colors)
	// updateMap(colors, "black", "#000000")
	// updateMap(colors, "ww", "#111111")
	// printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hexcode := range colors {
		fmt.Println(color, hexcode)
	}
}

func updateMap(colors map[string]string, color string, code string) {
	colors[color] = code
}
