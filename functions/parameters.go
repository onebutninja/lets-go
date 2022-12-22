package main

import (
	"fmt"
)


func main () {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 3)
	fmt.Println(values)

	val := 10
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)
}


func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}