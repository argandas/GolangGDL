package main

import "fmt"

func main() {
	animals := []string{"dog", "cat", "mouse"}
	for key, animal := range animals {
		fmt.Printf("[%d] %s\r\n", key, animal)
	}
}
