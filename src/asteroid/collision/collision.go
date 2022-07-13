package asteroid

func canCollide(a, b int) bool {
	return a > 0 && b < 0
}

func asteroidCollision(asteroids []int) []int {
	ret := []int{}
	for _, asteroid := range asteroids {
		for asteroid != 0 && len(ret) != 0 && canCollide(ret[len(ret)-1], asteroid) {
			temp := ret[len(ret)-1]
			if asteroid < 0 {
				if -asteroid >= temp {
					ret = ret[:len(ret)-1]
				}
				if -asteroid <= temp {
					asteroid = 0
				}
			} else {
				if asteroid > -temp {
					ret = ret[:len(ret)-1]
				}
				if asteroid <= -temp {
					asteroid = 0
				}
			}
		}
		if asteroid != 0 {
			ret = append(ret, asteroid)
		}
	}
	return ret
}
