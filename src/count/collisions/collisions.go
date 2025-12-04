package collisions

func countCollisions(directions string) int {
	var top [2]int
	var out int
	for _, direction := range directions {
		if direction == 'L' {
			if top[1] != 0 {
				if byte(top[0]) == 'R' {
					out += 1 + top[1]
					top[0], top[1] = int('S'), 1
				} else {
					out++
				}
			}
		} else if direction == 'R' {
			if top[1] != 0 {
				if byte(top[0]) == 'R' {
					top[1]++
				} else {
					top[0], top[1] = int('R'), 1
				}
			} else {
				top[0], top[1] = int('R'), 1
			}
		} else {
			if top[1] != 0 {
				if byte(top[0]) == 'R' {
					out += top[1]
					top[0], top[1] = int('S'), 1
				}
			} else {
				top[0], top[1] = int('S'), 1
			}
		}
	}
	return out
}
