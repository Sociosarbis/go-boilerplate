package three

func checkPowersOfThree(n int) bool {
	for n != 0 {
		r := n % 3
		if r == 2 {
			return false
		}
		if r == 1 {
			n -= 1
		} else {
			n /= 3
		}
	}
	return true
}
