package rover

import (
	"fmt"
)

func Execute(commands string) string {
	pos := Point{0, 0}
	directionIndex := 0

	for _, command := range commands {
		if command == 'R' {
			directionIndex++
			if directionIndex > 3 {
				directionIndex = 0
			}
		} else if command == 'L' {
			directionIndex--
			if directionIndex < 0 {
				directionIndex = 3
			}
		}

		direction := getDirection(directionIndex)

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
	return fmt.Sprintf("%d:%d:%s", pos.x, pos.y, getDirection(directionIndex))
}

func getDirection(directionIndex int) string {
	var m = map[int]string{0: "N", 1: "E", 2: "S", 3: "W"}
	return m[directionIndex]
}
