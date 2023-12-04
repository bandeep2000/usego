package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string
	Desc    string
}

func main1() {
	birdJson := `{"species": "foo", "desc": "food2", "species": "foo5", "desc": "food5"}`

	var b Bird

	json.Unmarshal([]byte(birdJson), &b)

	fmt.Println(b)

}

func GetMaxZeros(s1 string) int {

	maxCount := 0
	currCount := 0
	lastIndex := len(s1) - 1

	for i, v := range s1 {

		//fmt.Println(v)
		if v == '0' {

			currCount = currCount + 1

		}
		//This gets executed if we found 0 first time
		// and variable is not initialized

		// Get last index as 0011000 - last 0s strings also get executed

		if (v != '0') || (i == lastIndex) {
			if maxCount == 0 {
				maxCount = currCount
				currCount = 0
				continue

			}
			if maxCount < currCount {
				maxCount = currCount
				currCount = 0
			}
			//reset current count

		}

	}

	fmt.Println(maxCount)
	return maxCount

}

func main() {

	s1 := "0010"
	//s2 := strings.Fields(s1)

	//fmt.Println(s1)

	fmt.Println(s1)

}
