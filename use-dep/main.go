package main

import (
	"fmt"

	"github.com/bandeep2000/usego/greetings/fileutils"
	"github.com/bandeep2000/usego/greetings/hello"
)

func main() {

	fmt.Println(hello.Sayhello())
	file1 := fileutils.New("p2.txt")

	//fmt.Println(file)
	fmt.Println(file1.CheckForStringInFile("blah"))

}
