package fileutils

import (
	"testing"
)

// Tests if string is present
func TestFileContains(t *testing.T) {

	file := New("p2.txt")
	stringToChk := "together"
	isStrPresent, _ := file.CheckForStringInFile(stringToChk)

	if isStrPresent != true {

		t.Errorf("Test failed: function returned true although string '%v' was not found", stringToChk)

	}

}

// Tests file exists in log.Fatal won't allow to unit test with some go limitations, this will regardless
func TestFileExists(t *testing.T) {

	//file := fileDetails{"noexistingfile"}

	file := New("noexistingfile")

	_, err := file.CheckForStringInFile("together")

	if err == nil {

		t.Error("Test failed file does not exits")

	}

}
