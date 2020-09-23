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

		if command == 'M' {
			if direction == "E" {
				pos.moveEast()
			} else if direction == "W" {
				pos.moveWest()
			} else if direction == "S" {
				pos.moveSouth()
			} else {
				pos.moveNorth()
			}
		}

	}
	return fmt.Sprintf("%d:%d:%s", pos.x, pos.y, compass.getDirection())
}
