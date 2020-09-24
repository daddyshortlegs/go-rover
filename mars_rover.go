package rover

import (
	"fmt"
)

type Rover struct {
	pos     Point
	bearing Bearing
}

func Execute(commands string) string {
	pos := Point{0, 0}

	rover := Rover{}

	for _, command := range commands {
		if command == 'R' {
			turnRight(&rover)
		} else if command == 'L' {
			turnLeft(&rover)
		}

		if command == 'M' {
			move(&pos, rover.bearing.getDirection())
		}
	}
	return fmt.Sprintf("%d:%d:%s", pos.x, pos.y, rover.bearing.getDirection())
}

func turnLeft(rover *Rover) {
	rover.bearing.left()
}

func turnRight(rover *Rover) {
	rover.bearing.right()
}

func move(pos *Point, direction string) {
	m := make(map[string]func())
	m["E"] = pos.moveEast
	m["W"] = pos.moveWest
	m["N"] = pos.moveNorth
	m["S"] = pos.moveSouth
	m[direction]()
}
