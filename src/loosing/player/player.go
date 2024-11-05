package player

func losingPlayer(x int, y int) string {
	y /= 4
	var t int
	if x < y {
		t = x
	} else {
		t = y
	}
	if t%2 == 0 {
		return "Bob"
	} else {
		return "Alice"
	}
}
