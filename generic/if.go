// Conditionals
package main

import (
	"fmt"
)

func main() {
	x := 1

	if x > 5 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 0 && x < 2 {
		fmt.Println("x is just right")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}


	n := 5

	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		println("vague")
	}
}
