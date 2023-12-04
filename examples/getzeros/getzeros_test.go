package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeros(t *testing.T) {

	assert.Equal(t, GetMaxZeros("00000"), 5, "count doesn't match")
}
