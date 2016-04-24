package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDungeon(t *testing.T) {
	dungeon := NewDungeon(5, 10)

	assert.Equal(t, 5, len(dungeon.Tiles), "Width should be 5")
	assert.Equal(t, 10, len(dungeon.Tiles[0]), "Height should be 10")
}
