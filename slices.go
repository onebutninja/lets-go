package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "duffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)		

	fmt.Println((len(loons)))
	fmt.Println(loons[1:])

	for _, name := range loons {
		fmt.Println(name)
	}

	loons = append(loons, "elmer")
	fmt.Println(loons)
}
