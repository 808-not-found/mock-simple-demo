package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOutput(t *testing.T) {
	file := FileEater{Path: "text.txt"}
	assert.Equal(t, "世界\n你好\n", ReverseOutput(file))
}
