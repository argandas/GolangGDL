// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/argandas/repos")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	// var r Response
	r := Response{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range r {
		fmt.Println(repo.Name)
	}
}

type Response []Repo

type Repo struct {
	Name string
	URL   string
}
