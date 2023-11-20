# Go basic module

This repo provides basic usage of go with modules and packages

## Installation

Please refer to this to install go:
[go install ](https://go.dev/doc/install).

## Example usage

Note: Please refer to use-dep

```go
package main

import (
	"fmt"
    
	"github.com/bandeep2000/usego/greetings/fileutils"
	"github.com/bandeep2000/usego/greetings/hello"
)

func main() {
     // Just say hello using hello package
	fmt.Println(hello.Sayhello())

    // Check for file contains a string
   file1 := fileutils.New("p2.txt")
   // check for for error as it return nil
   fmt.Println(file1.CheckForStringInFile("blah"))

}
```
**Note:** Above example assumes you have *p2.txt* in the current directory with *blah* string to find in the file
## Run 

``` 
go mod init example 
go mod tidy
go run . 
```
Note: It might ask you to run *go get <module name>* 

Example output:
```
go run .
hello1
false <nil>
```
