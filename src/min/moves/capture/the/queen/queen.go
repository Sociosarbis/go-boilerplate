package queen

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func isSameSignSmall(a, b int) bool {
	if a*b < 0 {
		return false
	}
	return abs(a) < abs(b)
}

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if (a == e && !(c == e && isSameSignSmall(f-d, f-b))) || (b == f && !(d == f && isSameSignSmall(e-c, e-a))) {
		return 1
	}
	dx := e - c
	dy := f - d
	dx2 := e - a
	dy2 := f - b
	if abs(dx) == abs(dy) && !(abs(dx2) == abs(dy2) && isSameSignSmall(e-a, dx) && isSameSignSmall(f-b, dy)) {
		return 1
	}
	return 2
}
