package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.35,
		"GOOG": 2135.22,
		"MSFT": 1326.15,
	}

	fmt.Println(len(stocks), stocks)
	fmt.Println(stocks["MSFT"])

	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	
	stocks["TSLA"] = 811.22
	fmt.Println(stocks)

	delete(stocks, "AMZN")
	fmt.Println(stocks)

	for key, value := range stocks {
		fmt.Println(key, value)
	}
}