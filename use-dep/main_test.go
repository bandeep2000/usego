package main

import (
	"testing"

	"github.com/bandeep2000/usego/greetings/hello"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	assert.Equal(t, "hello1", hello.Sayhello(), "Output not matching")

}
