// +build OMIT

package main

import (
	"fmt"
	"github.com/argandas/githubapi" // HL
	"log"
)

func main() {
	items, err := githubapi.Get("argandas") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
