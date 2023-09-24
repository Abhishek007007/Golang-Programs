package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#32435",
	}

	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, code := range c {
		println(color, " : ", code)
	}
}
