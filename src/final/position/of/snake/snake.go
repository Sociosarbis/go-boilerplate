package snake

func finalPositionOfSnake(n int, commands []string) int {
	var ret int
	for _, command := range commands {
		switch command {
		case "RIGHT":
			ret++
		case "LEFT":
			ret--
		case "UP":
			ret -= n
		case "DOWN":
			ret += n
		}
	}
	return ret
}
