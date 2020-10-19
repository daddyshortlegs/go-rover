package rover

type Grid struct {
	maxX int
	maxY int
}

func (grid Grid) width() int {
	return grid.maxX
}

func (grid Grid) height() int {
	return grid.maxY
}

func (grid Grid) adjust(pos Point) Point {
	if pos.x < 0 {
		return Point{9, pos.y}
	}

	if pos.y < 0 {
		return Point{pos.x, 9}
	}

	return Point{
		x: pos.x % 10,
		y: pos.y % 10,
	}
}
