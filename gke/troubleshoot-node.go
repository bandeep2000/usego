package main

import (
	"fmt"
	"gke/troubleshooting"
)

func PrintArray(arr1 []string) {

	for _, v := range arr1 {

		fmt.Println(v)

	}

}

func main() {

	PrintArray(troubleshooting.GetNodeFailureCmds())

}
