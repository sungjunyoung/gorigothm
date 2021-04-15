package robot_bounded_in_circle

func isRobotBounded(instructions string) bool {
	x := 0
	y := 0
	direction := 0

	for _, instruction := range instructions {

		switch string(instruction) {
		case "G":
			switch direction % 4 {
			case 0:
				y += 1
			case 1, -1:
				x += 1
			case 2, -2:
				y -= 1
			case 3, -3:
				x -= 1
			}
		case "R":
			direction += 1
		case "L":
			direction -= 1
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	return direction%4 != 0
}
