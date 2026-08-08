package number2

func getProduct(n int) int {
	out := 1
	for n != 0 {
		out *= n % 10
		n /= 10
	}
	return out
}

func smallestNumber(n int, t int) int {
	for {
		if getProduct(n)%t == 0 {
			return n
		}
		n++
	}
}
