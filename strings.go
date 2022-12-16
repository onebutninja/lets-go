// GO Strings

package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println("t" + book[1:])
	
	poem := `
		San vgeis ston phgaimo gia
		thn Ithaki na euhesai
		na einai makris o dromos
	`
	fmt.Println(poem)
}	