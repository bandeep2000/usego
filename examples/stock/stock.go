package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const financeAPI = "https://query1.finance.yahoo.com/v8/finance/chart/"

//type Quote struct {
//	RegularMarketPrice float64 `json:"regularMarketPrice"`
//}

func get_quote(symbol string) (float64, error) {

	url := fmt.Sprintf(financeAPI + symbol)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching stock quote:", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//TODO: Create unit tests to test this
		fmt.Println("Error reading response body:", err)
		return 0, err
	}

	/*

			(map[string]interface{}) because the structure of the JSON response might contain nested objects and
			arrays with varying keys and values.
			Using map[string]interface{} allows us to handle this
		    dynamic structure without defining a specific struct beforehand.

	*/
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0, err
	}

	// Check for error , the above error checking did not work, as api will return
	// correctly even if invalid stock is passed
	error_stock := (data["chart"].(map[string]interface{})["error"])
	if error_stock != nil {
		//log.Fatal("error getting stock")
		return 0, errors.New("error getting stock")
	}

	quoteData := data["chart"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["regularMarketPrice"].(float64)
	return quoteData, nil

}
