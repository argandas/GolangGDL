// +build OMIT

package main

import (
	"fmt"
	"github.com/argandas/wordpress" // HL
	"log"
)

func main() {
	items, err := wordpress.Get("en.blog.wordpress.com") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
