package rover

import "fmt"

type Rover struct {
	pos     Point
	bearing Bearing
	grid    Grid
}

func (rover *Rover) turnLeft() {
	rover.bearing.left()
}

func (rover *Rover) turnRight() {
	rover.bearing.right()
}

func (rover *Rover) move() {
	m := make(map[string]func())
	m["E"] = rover.moveEast
	m["W"] = rover.moveWest
	m["N"] = rover.moveNorth
	m["S"] = rover.moveSouth
	m[rover.bearing.getDirection()]()
}

func (rover *Rover) moveEast() {
	rover.pos = rover.pos.Add(Point{1, 0})
	if rover.grid.outOfBounds(rover.pos) {
		rover.pos = Point{0, rover.pos.y}
	}
}

func (rover *Rover) moveWest() {
	rover.pos = rover.pos.Add(Point{-1, 0})
	if rover.grid.outOfBounds(rover.pos) {
		rover.pos = Point{9, rover.pos.y}
	}
}

func (rover *Rover) moveSouth() {
	rover.pos = rover.pos.Add(Point{0, -1})
	if rover.grid.outOfBounds(rover.pos) {
		rover.pos = Point{rover.pos.x, 9}
	}
}

func (rover *Rover) moveNorth() {
	rover.pos = rover.pos.Add(Point{0, 1})
	if rover.grid.outOfBounds(rover.pos) {
		rover.pos = Point{rover.pos.x, 0}
	}
}

func (rover *Rover) String() string {
	return fmt.Sprintf("%d:%d:%s", rover.pos.x, rover.pos.y, rover.bearing.getDirection())
}
