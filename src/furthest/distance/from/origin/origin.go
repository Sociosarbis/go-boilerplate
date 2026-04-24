package origin

func furthestDistanceFromOrigin(moves string) int {
	var temp, out int
	for _, c := range moves {
		if c == 'L' || c == '_' {
			temp++
		} else {
			temp--
		}
	}
	if temp > out {
		out = temp
	}
	temp = 0
	for _, c := range moves {
		if c == 'R' || c == '_' {
			temp++
		} else {
			temp--
		}
	}
	if temp > out {
		out = temp
	}
	return out
}
