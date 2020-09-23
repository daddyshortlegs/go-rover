package rover

import (
	"fmt"
)

type Compass struct {
	directionIndex int
}

func Execute(commands string) string {
	pos := Point{0, 0}

	compass := Compass{0}

	for _, command := range commands {
		if command == 'R' {
			compass.directionIndex++
			if compass.directionIndex > 3 {
				compass.directionIndex = 0
			}
		} else if command == 'L' {
			compass.directionIndex--
			if compass.directionIndex < 0 {
				compass.directionIndex = 3
			}
		}

		direction := getDirection(compass.directionIndex)

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
	return fmt.Sprintf("%d:%d:%s", pos.x, pos.y, getDirection(compass.directionIndex))
}

func getDirection(directionIndex int) string {
	var m = map[int]string{0: "N", 1: "E", 2: "S", 3: "W"}
	return m[directionIndex]
}
