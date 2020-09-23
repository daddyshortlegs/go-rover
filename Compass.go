package rover

type Compass struct {
	directionIndex int
}

func (compass *Compass) left() {
	compass.directionIndex--
	if compass.directionIndex < 0 {
		compass.directionIndex = 3
	}
}

func (compass *Compass) right() {
	compass.directionIndex++
	if compass.directionIndex > 3 {
		compass.directionIndex = 0
	}
}

func (compass *Compass) getDirection() string {
	var m = map[int]string{0: "N", 1: "E", 2: "S", 3: "W"}
	return m[compass.directionIndex]
}
