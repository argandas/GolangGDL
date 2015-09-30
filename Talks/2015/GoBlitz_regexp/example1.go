package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MatchString("p([a-z]+)ch", "peaches"))
}
