package stair

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func n(f1, t1, t2 int) int {
	r1 := 1
	r2 := 1
	if f1 < 2 {
		f1 = 2
	}
	f2 := 2
	for f1 <= t1 || f2 <= t2 {
		if f1 <= t1 {
			r1 *= f1
		}
		if f2 <= t2 {
			r2 *= f2
		}
		g := gcd(r1, r2)
		if g != 1 {
			r1 /= g
			r2 /= g
		}
	}
	return r1 / r2
}

func waysToReachStair(k int) int {
	var x int
	base := 1
	var ret int
	for base-k <= x+1 {
		y := base - k
		if y >= 0 {
			ret += n(x+1-y, x+1, 2, y)
		}
		base *= 2
		x++
	}
	return ret
}
