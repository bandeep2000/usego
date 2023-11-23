package main

// Run as
// go run . AAPL # for AAPL symblo

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Command-line flag to specify the stock symbol
	if len(os.Args) < 2 {

		log.Fatal("Please pass stock price symbol")
	}

	symbol := os.Args[1]

	quoteData, err := get_quote(symbol)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("Stock Price for %s: %.2f\n", symbol, quoteData)
}
