package main

import (
	"bytes"
)

type Dungeon struct {
	Tiles [][]*Tile
}

func NewDungeon(width int, height int) *Dungeon {
	columns := make([][]*Tile, width)
	for x, _ := range columns {
		columns[x] = make([]*Tile, height)

		for y, _ := range columns[x] {
			columns[x][y] = NewTile("wall")
		}
	}

	return &Dungeon{
		Tiles: columns,
	}
}

func (d *Dungeon) Generate() {
}

func (d *Dungeon) String() string {
	buffer := bytes.Buffer{}

	for y := 0; y < d.height(); y++ {
		for x := 0; x < d.width(); x++ {
			tile := d.Tiles[x][y]

			buffer.WriteString(tile.String())
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (d *Dungeon) width() int {
	return len(d.Tiles)
}

func (d *Dungeon) height() int {
	return len(d.Tiles[0])
}
