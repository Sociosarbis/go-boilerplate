package number

func lcm(a int, b int) int64 {
	n1, n2 := int64(a), int64(b)
	if n2 > n1 {
		n1, n2 = n2, n1
	}
	for n1%n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return int64(a) / n2 * int64(b)
}

func nthMagicalNumber(n int, a int, b int) int {
	c := lcm(a, b)
	n1, n2 := int64(a), int64(b)
	nx := int64(n)
	var l int64
	var r int64
	if n1 < n2 {
		l = n1
		r = n1 * nx
	} else {
		l = n2
		r = n2 * nx
	}
	for l <= r {
		mid := (l + r) / 2
		count := mid/n1 + mid/n2 - mid/c
		if count >= nx {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return int(l % (1e9 + 7))
}
