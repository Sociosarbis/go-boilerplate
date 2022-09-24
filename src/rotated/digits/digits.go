package digits

func isValid(n int) bool {
	ret := false
	for n != 0 {
		temp := n % 10
		if temp == 2 || temp == 5 || temp == 6 || temp == 9 {
			if !ret {
				ret = true
			}
		} else if !(temp == 0 || temp == 1 || temp == 8) {
			return false
		}
		n /= 10
	}
	return ret
}

func rotatedDigits(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		if isValid(i) {
			ret++
		}
	}
	return ret
}
