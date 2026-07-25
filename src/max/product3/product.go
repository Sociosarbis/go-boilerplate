package product3

func top2(a, b, c int) (int, int) {
	if c > a {
		return c, a
	}
	if c > b {
		return a, c
	}
	return a, b
}

func maxProduct(n int) int {
	var a, b int
	for n != 0 {
		a, b = top2(a, b, n%10)
		n /= 10
	}
	return a * b
}
