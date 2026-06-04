package waviness

func isWave(digits []int, i int) bool {
	p, n := (i+2)%3, (i+1)%3
	return (digits[i] > digits[p] && digits[i] > digits[n]) || (digits[i] < digits[p] && digits[i] < digits[n])
}

func count(num int, digits []int) int {
	if num < 100 {
		return 0
	}
	for i := 0; i < 3; i++ {
		digits[i] = num % 10
		num /= 10
	}
	i := 1
	var out int
	if isWave(digits, i%3) {
		out++
	}
	i++
	for num != 0 {
		digits[(i+1)%3] = num % 10
		if isWave(digits, i%3) {
			out++
		}
		num /= 10
		i++
	}
	return out
}

func totalWaviness(num1 int, num2 int) int {
	var out int
	digits := make([]int, 3)
	for i := num1; i <= num2; i++ {
		out += count(i, digits)
	}
	return out
}
