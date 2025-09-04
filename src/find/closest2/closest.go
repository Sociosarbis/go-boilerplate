package closest2

func findClosest(x int, y int, z int) int {
	a := x - z
	if a < 0 {
		a = -a
	}
	b := y - z
	if b < 0 {
		b = -b
	}
	if a > b {
		return 2
	} else if a < b {
		return 1
	}
	return 0
}
