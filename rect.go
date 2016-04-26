package main

type Rect struct {
	X, Y, Width, Height int
}

func NewRect(x, y, width, height int) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r *Rect) Overlaps(other *Rect) bool {
	return r.X1() <= other.X2() && r.X2() >= other.X1() && r.Y1() <= other.Y2() && r.Y2() >= other.Y1()
}

func (r *Rect) Points() []*Point {
	points := []*Point{}
	for x := r.X1(); x <= r.X2(); x++ {
		for y := r.Y1(); y <= r.Y2(); y++ {
			points = append(points, &Point{x, y})
		}
	}
	return points
}

func (r *Rect) X1() int {
	return r.X
}

func (r *Rect) X2() int {
	return r.X + r.Width - 1
}

func (r *Rect) Y1() int {
	return r.Y
}

func (r *Rect) Y2() int {
	return r.Y + r.Height - 1
}
