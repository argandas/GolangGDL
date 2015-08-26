package main

import "fmt"

func main() {
	elements := map[string]string{
		"Hg": "Mercury",
		"Ag": "Silver",
		"Au": "Gold",
	}
	for symbol, element := range elements {
		fmt.Printf("[%s] %s\r\n", symbol, element)
	}
}
