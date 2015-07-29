// Package githubapi implements a basic client for the Github API.
// +build OMIT

package githubapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	items, err := Get("argandas") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items { // HL
		fmt.Println(item)
	}
}

type response []Repo

// Repo describes a Github repository.
type Repo struct {
	Name string
	Description string
	URL   string
	Stars int `json:"stargazers_count"`
}

func (r Repo) String() string {
	numStars := ""
	switch r.Stars {
	case 0:
		// nothing
	case 1:
		numStars = " (1 star)"
	default:
		numStars = fmt.Sprintf(" (%d stars)", r.Stars)
	}
	return fmt.Sprintf("\r\n%s\r\n%s%s\n[%s]", r.Name, r.Description, numStars, r.URL)
}

// Get fetches the repositories from a Github user.
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
	r := response{}     // HLdecode
	err = json.NewDecoder(resp.Body).Decode(&r) // HLdecode
	if err != nil {
		return nil, err // HLreturn
	}
	return r, nil // HLprepare
}