package rover

import (
	"fmt"
)

func Execute(commands string) string {
	pos := Point{0, 0}

	compass := Compass{0}

	for _, command := range commands {
		if command == 'R' {
			compass.right()
		} else if command == 'L' {
			compass.left()
		}

		direction := compass.getDirection()

		m := makeDirectionMap(pos)

		if command == 'M' {
			m[direction]()
		}

	}
	return fmt.Sprintf("%d:%d:%s", pos.x, pos.y, compass.getDirection())
}

func makeDirectionMap(pos Point) map[string]func() {
	m := make(map[string]func())
	m["E"] = pos.moveEast
	m["W"] = pos.moveWest
	m["N"] = pos.moveNorth
	m["S"] = pos.moveSouth
	return m
}
