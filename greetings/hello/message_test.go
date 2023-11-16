package hello

import (
	"testing"
)

func TestHello(t *testing.T) {

	if Sayhello() != "hello" {

		t.Errorf("Found error")

	}

}
