package main

const WallType = "wall"
const WallTypeChar = "x"
const DefaultTypeChar = " "

type Tile struct {
	Type string
}

func NewTile(t string) *Tile {
	return &Tile{
		Type: t,
	}
}

func (t *Tile) String() string {
	switch {
	case t.Type == WallType:
		return WallTypeChar
	default:
		return DefaultTypeChar
	}

}
