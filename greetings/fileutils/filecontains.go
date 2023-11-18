package fileutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func (f fileDetails) CheckForStringInFile(stringToChk string) (bool, error) {

	file, err := os.Open(f.filePath)
	if err != nil {
		log.Fatal(err)
		fmt.Println("executed")
		return false, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		if strings.Contains(scanner.Text(), stringToChk) {
			fmt.Println("Found string Contains")
			return true, nil

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return false, err
	}

	return false, nil

}

type fileDetails struct {
	filePath string
}

// Constructor, returns pointer to the file structures
// Teaks inputs with file name
func New(filename string) fileDetails {

	file := fileDetails{filename}

	return file

}
