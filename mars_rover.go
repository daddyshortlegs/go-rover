package rover

func Execute(commands string) string {
	grid := Grid{maxX: 9, maxY: 9}
	var rover = Rover{grid: grid}

	for _, command := range commands {
		run(&rover, command)
	}
	return rover.String()
}

func run(rover *Rover, command int32) {
	createCommands(rover)[command]()
}

func createCommands(rover *Rover) map[int32]func() {
	m := make(map[int32]func())
	m['R'] = rover.turnRight
	m['L'] = rover.turnLeft
	m['M'] = rover.move
	return m
}
