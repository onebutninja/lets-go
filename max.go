package main

import (
	"fmt"
)

func main() {
	nums := []int{88, 16, 8, 42, 3, 23, 15, 67}
	fmt.Println(nums)

	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}

	} 

	fmt.Println(max)
}