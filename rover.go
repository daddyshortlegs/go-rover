package rover

import (
	"fmt"
)

func Execute(commands string) string {
	y := 0

	for _, command := range commands {
		if command == 'M' {
			y++
		}

	}
	return fmt.Sprintf("0:%d:N", y)
}
