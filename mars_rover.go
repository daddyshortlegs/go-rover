package rover

import (
	"fmt"
)

func Execute(commands string) string {
	rover := Rover{}

	for _, command := range commands {
		if command == 'R' {
			rover.turnRight()
		} else if command == 'L' {
			rover.turnLeft()
		} else if command == 'M' {
			rover.move(rover.bearing.getDirection())
		}
	}
	return fmt.Sprintf("%d:%d:%s", rover.pos.x, rover.pos.y, rover.bearing.getDirection())
}
