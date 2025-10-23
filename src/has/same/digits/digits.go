package digits

func hasSameDigits(s string) bool {
	n := len(s)
	digits := make([]byte, n)
	for i, c := range s {
		digits[i] = byte(c - 48)
	}
	for i := n - 2; i > 0; i-- {
		for j := 0; j < i+1; j++ {
			digits[j] = (digits[j] + digits[j+1]) % 10
		}
	}
	return digits[0] == digits[1]
}
