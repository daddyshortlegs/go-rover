package rover

func Execute(commands string) string {
	rover := Rover{}

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
