package divisibility

func sum(n int) int {
	var a int
	b := 1
	for n != 0 {
		r := n % 10
		a += r
		b *= r
		n /= 10
	}
	return a + b
}

func checkDivisibility(n int) bool {
	return n%sum(n) == 0
}
