package rover

func Execute(commands string) string {
	grid := Grid{maxX: 9, maxY: 9}
	var rover = Rover{grid: grid}
	for _, command := range commands {
		if command == 'R' {
			rover.turnRight()
		} else if command == 'L' {
			rover.turnLeft()
		} else if command == 'M' {
			rover.move()
		}
	}
	return rover.String()
}
