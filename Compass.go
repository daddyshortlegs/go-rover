package rover

type Bearing struct {
	directionIndex int
}

func (bearing *Bearing) left() {
	bearing.directionIndex--
	if bearing.directionIndex < 0 {
		bearing.directionIndex = 3
	}
}

func (bearing *Bearing) right() {
	bearing.directionIndex++
	if bearing.directionIndex > 3 {
		bearing.directionIndex = 0
	}
}

func (bearing *Bearing) getDirection() string {
	var m = map[int]string{0: "N", 1: "E", 2: "S", 3: "W"}
	return m[bearing.directionIndex]
}
