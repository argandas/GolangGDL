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
	items, err := Get("en.blog.wordpress.com") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items { // HL
		fmt.Println(item)
	}
}

type Response struct {
	Posts []Post
}

type Post struct {
	Title string
	URL   string
	Likes int `json:"like_count"`
}

func (p Post) String() string {
	numLikes := ""
	switch p.Likes {
	case 0:
		// nothing
	case 1:
		numLikes = " (1 like)"
	default:
		numLikes = fmt.Sprintf(" (%d likes)", p.Likes)
	}
	return fmt.Sprintf("\r\n%s%s\n[%s]", p.Title, numLikes, p.URL)
}

func Get(wp string) ([]Post, error) {
	url := fmt.Sprintf("https://public-api.wordpress.com/rest/v1.1/sites/%s/posts/?number=10", wp) // HLurl
	resp, err := http.Get(url)                                                                     // HLget
	if err != nil {
		return nil, err // HLreturn
	}
	defer resp.Body.Close()               // HLclose
	if resp.StatusCode != http.StatusOK { // HLstatus
		return nil, errors.New(resp.Status) // HLerrors
	}
	r := new(Response)                         // HLdecode
	err = json.NewDecoder(resp.Body).Decode(r) // HLdecode
	if err != nil {
		return nil, err // HLreturn
	}
	return r.Posts, nil // HLreturn
}
