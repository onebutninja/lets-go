package main

import (
	"fmt"
	"strings"
)

func main() {

	text := `
		my name is same as my Name
		and my NaMe is Dimitrios Oikonomidis,
		not just dimitRios
		`

	words := strings.Fields(text)
	counts := map[string]int{}

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)

}
