package integers

func hasZeros(n int) bool {
	for n != 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}

func getNoZeroIntegers(n int) []int {
	for i := n / 2; i > 0; i-- {
		if !hasZeros(i) && !hasZeros(n-i) {
			return []int{i, n - i}
		}
	}
	return nil
}
