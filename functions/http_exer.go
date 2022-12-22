package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // Make sure you close the body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // return error if Ctype haeder not found
		return "", fmt.Errorf("Cannot find Ctype header")
	}

	return ctype, nil
	
}
