package main

const WallType = "wall"
const WallTypeChar = "xx"
const FloorType = "floor"
const FloorTypeChar = "  "
const DefaultTypeChar = "!!"

type Tile struct {
	Point
	Type   string
	region int
}

func NewTile(t string, x, y int) *Tile {
	return &Tile{
		Type:  t,
		Point: Point{x, y},
	}
}

func (t *Tile) String() string {
	return map[string]string{
		WallType:  WallTypeChar,
		FloorType: FloorTypeChar,
	}[t.Type]
}
