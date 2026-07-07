package multiply

func sumAndMultiply(n int) int64 {
	var sum, x int
	base := 1
	for n != 0 {
		if n%10 != 0 {
			x += base * (n % 10)
			sum += n % 10
			if base < 1e9 {
				base *= 10
			}
		}
		n /= 10
	}
	return int64(x) * int64(sum)
}
