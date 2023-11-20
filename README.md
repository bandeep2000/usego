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

   fmt.Println(file1.CheckForStringInFile("blah"))

}
```

