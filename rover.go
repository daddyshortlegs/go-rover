package rover

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func Execute(commands string) string {
	pos := Point{0, 0}
	directionIndex := 0
	direction := "N"

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

		if directionIndex == 0 {
			direction = "N"
		} else if directionIndex == 1 {
			direction = "E"
		} else if directionIndex == 2 {
			direction = "S"
		} else if directionIndex == 3 {
			direction = "W"
		}

		if command == 'M' {
			if direction == "E" {
				pos.x++
				if pos.x > 9 {
					pos.x = 0
				}
			} else if direction == "W" {
				pos.x--
				if pos.x < 0 {
					pos.x = 9
				}
			} else if direction == "S" {
				pos.y--
				if pos.y < 0 {
					pos.y = 9
				}
			} else {
				pos.y++
				if pos.y > 9 {
					pos.y = 0
				}
			}
		}

	}
	return fmt.Sprintf("%d:%d:%s", pos.x, pos.y, direction)
}
