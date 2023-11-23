package main

import (
	"fmt"
	"testing"
)

func TestStock(t *testing.T) {

	price, _ := get_quote("AAPL")

	// Additional test to check if the returned value is of type float64
	if _, ok := interface{}(price).(float64); !ok {
		t.Errorf("Returned value is not of type float64")
	}
}

func TestInvalidStock(t *testing.T) {

	_, err := get_quote("AAPL1")

	// Additional test to check if the returned value is of type float64
	if err == nil {
		//t.Errorf("Expected error as nil but ")
		fmt.Println(err)
		t.Errorf("Expected price error as nil")
	}
}
