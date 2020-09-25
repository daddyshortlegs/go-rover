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

func (pos *Point) moveEast() {
	*pos = pos.Add(Point{1, 0})
	//pos.x++
	if pos.x > 9 {
		pos.x = 0
	}
}

func (pos *Point) moveWest() {
	pos.x--
	if pos.x < 0 {
		pos.x = 9
	}
}

func (pos *Point) moveSouth() {
	pos.y--
	if pos.y < 0 {
		pos.y = 9
	}
}

func (pos *Point) moveNorth() {
	pos.y++
	if pos.y > 9 {
		pos.y = 0
	}
}
