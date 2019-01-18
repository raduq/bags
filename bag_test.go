package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	b := New(2)
	assert.Equal(t, 2, b.Size, "Size must be 2")
}
