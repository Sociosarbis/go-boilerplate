package permutation

func period(i int, halfN int) int {
	ii := i
	ret := 0
	for {
		if i&1 == 1 {
			i = halfN + (i-1)/2
		} else {
			i /= 2
		}
		ret++
		if i == ii {
			break
		}
	}
	return ret
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func reinitializePermutation(n int) int {
	ret := 1
	halfN := n / 2
	for i := 2; i < n; i++ {
		p := period(i, halfN)
		g := gcd(p, ret)
		ret = p * ret / g
	}
	return ret
}
