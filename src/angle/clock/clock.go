package clock

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func angleClock(hour int, minutes int) float64 {
	hour %= 12
	a := float64(minutes * 6)
	b := float64(hour*30) + float64(minutes)*0.5
	if a > b {
		a, b = b, a
	}
	return min(abs(a-b), abs(a+360-b))
}
