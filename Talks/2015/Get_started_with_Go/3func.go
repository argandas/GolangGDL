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
	posts, err := Get("en.blog.wordpress.com") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts { // HL
		fmt.Println(post.Title)
	}
}

type Post struct {
	Title string
	URL   string
}

type Response struct {
	Posts []Post
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
	return r.Posts, nil // HLprepare
}
