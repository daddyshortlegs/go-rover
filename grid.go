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

func (grid Grid) outOfBounds(pos Point) bool {
	return pos.x > grid.maxX || pos.x < 0 || pos.y < 0 || pos.y > grid.maxY
}
