// +build OMIT

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repos, err := Get("argandas") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range repos { // HL
		fmt.Println(repo.Name)
	}
}

type Response []Repo

type Repo struct {
	Name string
	URL   string
}

func Get(user string) ([]Repo, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", user) // HLurl
	resp, err := http.Get(url)   // HLget
	if err != nil {
		return nil, err // HLreturn
	}
	defer resp.Body.Close()               // HLclose
	if resp.StatusCode != http.StatusOK { // HLstatus
		return nil, errors.New(resp.Status) // HLerrors
	}
	r := Response{}     // HLdecode
	err = json.NewDecoder(resp.Body).Decode(&r) // HLdecode
	if err != nil {
		return nil, err // HLreturn
	}
	return r, nil // HLprepare
}
