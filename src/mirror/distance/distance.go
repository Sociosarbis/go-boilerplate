package distance

func reverse(num int) int {
	var out int
	for num != 0 {
		d := num % 10
		out = out*10 + d
		num /= 10
	}
	return out
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func mirrorDistance(n int) int {
	return abs(n - reverse(n))
}
