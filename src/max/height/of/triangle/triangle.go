package triangle

func needAmount(n int) (int, int) {
	if n%2 == 0 {
		return (2 + n) * n / 4, n * n / 4
	} else {
		return (1 + n) * (n + 1) / 4, (n + 1) * (n / 2) / 2
	}
}

func maxHeightOfTriangle(red int, blue int) int {
	l := 1
	var r int
	if red > blue {
		r = red
	} else {
		r = blue
	}
	ret := l
	for l <= r {
		mid := (l + r) / 2
		a, b := needAmount(mid)
		if (red >= a && blue >= b) || (red >= b && blue >= a) {
			if mid > ret {
				ret = mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
