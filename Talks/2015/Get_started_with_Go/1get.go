// +build OMIT

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() { // HLfunc
	resp, err := http.Get("https://api.github.com/users/argandas/repos") // HLget
	if err != nil {                                                                                                  // HLerr
		log.Fatal(err) // HLerr
	} // HLerr
	if resp.StatusCode != http.StatusOK { // HLstatus
		log.Fatal(resp.Status) // HLstatus
	} // HLstatus
	_, err = io.Copy(os.Stdout, resp.Body) // HLcopy
	if err != nil {                        // HLerr
		log.Fatal(err) // HLerr
	} // HLerr
} // HLfunc
