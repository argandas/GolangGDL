package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r.ReplaceAllString("a peach pinch", "<fruit>"))
}
