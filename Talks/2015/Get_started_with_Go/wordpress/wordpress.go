// Package wordpress implements a basic client for the Wordpress API.
// +build OMIT

package wordpress

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Post describes a Wordpress post.
type Post struct {
	Title string
	URL   string
	Likes int `json:"like_count"`
}

func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		// nothing
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

// Get fetches the most recent posts from a Wordpress blog.
func Get(wp string) ([]Post, error) {
	url := fmt.Sprintf("https://public-api.wordpress.com/rest/v1.1/sites/%s/posts/?number=10", wp) // HLurl
	resp, err := http.Get(url)                                // HLget
	if err != nil {
		return nil, err // HLreturn
	}
	defer resp.Body.Close()               // HLclose
	if resp.StatusCode != http.StatusOK { // HLstatus
		return nil, errors.New(resp.Status) // HLerrors
	}
	r := new(response)                         // HLdecode
	err = json.NewDecoder(resp.Body).Decode(r) // HLdecode
	if err != nil {
		return nil, err // HLreturn
	}
	return r.Posts, nil // HLreturn
}

type response struct {
	Posts []Post
}