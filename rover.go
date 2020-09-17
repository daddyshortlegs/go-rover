package rover

import (
	"fmt"
)

func Execute(commands string) string {
	x := 0
	y := 0
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
				x++
				if x > 9 {
					x = 0
				}
			} else if direction == "W" {
				x--
				if x < 0 {
					x = 9
				}
			} else if direction == "S" {
				y--
				if y < 0 {
					y = 9
				}
			} else {

				y++
				if y > 9 {
					y = 0
				}
			}
		}

	}
	return fmt.Sprintf("%d:%d:%s", x, y, direction)
}
