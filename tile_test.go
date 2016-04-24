package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTile(t *testing.T) {
	tile := NewTile("wall")

	assert.Equal(t, "x", tile.String(), "String should be x")
}
