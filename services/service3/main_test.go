package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var a string = "Ahoj"
	var b string = "Ahoj"

	assert.Equal(t, a, b, "two words should be the same")

	fmt.Println("unit test pass")
}
