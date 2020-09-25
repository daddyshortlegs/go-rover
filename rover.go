package rover

type Rover struct {
	pos     Point
	bearing Bearing
}

func (rover *Rover) turnLeft() {
	rover.bearing.left()
}

func (rover *Rover) turnRight() {
	rover.bearing.right()
}

func (rover *Rover) move(direction string) {
	m := make(map[string]func())
	m["E"] = rover.pos.moveEast
	m["W"] = rover.pos.moveWest
	m["N"] = rover.pos.moveNorth
	m["S"] = rover.pos.moveSouth
	m[direction]()
}
