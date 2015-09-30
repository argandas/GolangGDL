package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.MatchString("peach"))
}
