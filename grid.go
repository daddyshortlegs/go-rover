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
