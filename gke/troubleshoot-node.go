package main

import (
	"fmt"
	"gke/etcdcmd"
	"gke/troubleshooting"
	"os"
)

func PrintArray(arr1 []string) {

	for _, v := range arr1 {

		fmt.Println(v)

	}

}

func main() {

	if len(os.Args) == 1 {

		fmt.Println("Please pass an argument")
		os.Exit(1)

	}

	if os.Args[1] == "t" {

		PrintArray(troubleshooting.GetNodeFailureCmds())

	}

	if os.Args[1] == "e" {

		PrintArray(etcdcmd.GetCmds())

	}

}
