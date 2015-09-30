package main

import (
	"fmt"
	"regexp"
)

var r = regexp.MustCompile("p([a-z]+)ch")

func main() {
	fmt.Println(r.MatchString("peach"))
}
