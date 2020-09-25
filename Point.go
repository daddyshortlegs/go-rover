package rover

type Point struct {
	x int
	y int
}

func (pos Point) Add(other Point) Point {
	return Point{
		x: pos.x + other.x,
		y: pos.y + other.y,
	}
}
