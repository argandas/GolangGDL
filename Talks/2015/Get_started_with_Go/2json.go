// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://public-api.wordpress.com/rest/v1.1/sites/en.blog.wordpress.com/posts/?number=10")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range r.Posts {
		fmt.Println(post.Title)
	}
}

type Response struct {
	Posts []Post
}

type Post struct {
	Title string
	URL   string
}
