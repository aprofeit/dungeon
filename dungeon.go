package main

import (
	"bytes"
	"math/rand"
)

type Dungeon struct {
	Tiles         [][]*Tile
	Rooms         []*Rect
	currentRegion int
}

func NewDungeon(width int, height int) *Dungeon {
	tiles := make([][]*Tile, width)
	for x, _ := range tiles {
		tiles[x] = make([]*Tile, height)

		for y, _ := range tiles[x] {
			tiles[x][y] = NewTile(WallType, x, y)
		}
	}

	return &Dungeon{
		Tiles:         tiles,
		currentRegion: 0,
	}
}

func (d *Dungeon) Generate() {
	d.addRooms()

	for x := 1; x < d.width(); x += 2 {
		for y := 1; y < d.height(); y += 2 {
			tile := d.Tiles[x][y]

			if tile.Type == WallType {
				continue
			}

			d.growMaze(x, y)
		}
	}
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

func (d *Dungeon) CarveRect(r *Rect, rectType string) {
	for _, point := range r.Points() {
		tile := d.Tiles[point.X][point.Y]
		tile.Type = rectType
		tile.region = d.currentRegion
	}
}

const MAX_ROOM_ATTEMPTS = 1000
const ROOM_SIZE = 1

func (d *Dungeon) addRooms() {
	for i := 0; i < MAX_ROOM_ATTEMPTS; i++ {
		size := ((rand.Intn(3+ROOM_SIZE) + 1) * 2) + 1
		rectangularity := rand.Intn(1+size/2) * 2
		width := size
		height := size

		if rand.Intn(2) == 1 {
			width += rectangularity
		} else {
			height += rectangularity
		}

		x := rand.Intn((d.width()-width)/2)*2 + 1
		y := rand.Intn((d.height()-height)/2)*2 + 1

		room := NewRect(x, y, width, height)
		overlaps := false
		for _, otherRoom := range d.Rooms {
			if room.Overlaps(otherRoom) {
				overlaps = true
				break
			}
		}

		if overlaps {
			continue
		}

		d.Rooms = append(d.Rooms, room)
		d.currentRegion += 1
		d.CarveRect(room, FloorType)
	}
}

func (d *Dungeon) growMaze(x int, y int) {
}

func (d *Dungeon) width() int {
	return len(d.Tiles)
}

func (d *Dungeon) height() int {
	return len(d.Tiles[0])
}
