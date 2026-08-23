package game

func getA(value int) int {
	if value%2 == 0 {
		return value / 2
	}
	return value/2 + 1
}

func sumGame(num string) bool {
	var lc, rc, value int
	n := len(num)
	for i := n/2 - 1; i >= 0; i-- {
		if num[i] == '?' {
			lc++
		} else {
			value += int(num[i] - '0')
		}
	}
	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			rc++
		} else {
			value -= int(num[i] - '0')
		}
	}
	if lc < rc {
		rc -= lc
		lc = 0
	} else {
		lc -= rc
		rc = 0
	}
	if lc == 0 {
		if rc == 0 {
			return value != 0
		} else {
			if value > 0 {
				a := getA(rc)
				if value < 9*a {
					return true
				}
				ob := -9 * (rc - a)
				if ob+value > 0 {
					return true
				}
			} else {
				return true
			}
		}
	} else {
		if value >= 0 {
			return true
		} else {
			a := getA(lc)
			if value > -9*a {
				return true
			}
			ob := 9 * (lc - a)
			if ob+value < 0 {
				return true
			}
		}
	}
	return false
}
