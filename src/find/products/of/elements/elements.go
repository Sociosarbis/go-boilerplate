package elements

func findProductsOfElements(queries [][]int64) []int {
	ret := make([]int, len(queries))
	for i, q := range queries {
		from, to, mod := q[0], q[1], q[2]
		ret[i] = modPow(2, (getValue(to+1) - getValue(from)), mod)
	}
	return ret
}

func modPow(x int64, n int64, mod int64) int {
	if n == 0 {
		return int(1 % mod)
	}
	if n%2 == 1 {
		return int(x * int64(modPow(x, n-1, mod)) % mod)
	}
	return modPow(x*x%mod, n/2, mod)
}

func countOnes2(n int64) int64 {
	if n < 1 {
		return 0
	}
	return n * (1 << (n - 1))
}

func countOnes(n int64) int64 {
	if n <= 0 {
		return 0
	}
	o := n
	i := 0
	for n != 0 {
		n >>= 1
		i++
	}
	remain := o - (1 << (i - 1))
	return countOnes2(int64(i-1)) + countOnes(remain) + remain + 1
}

func countPowers2(n int64) int64 {
	if n <= 1 {
		return 0
	}
	return n * (n - 1) / 2 * (1 << (n - 1))
}

func countPowers(n int64) int64 {
	if n <= 1 {
		return 0
	}
	o := n
	i := 0
	for n != 0 {
		n >>= 1
		i++
	}
	remain := o - (1 << (i - 1))
	return countPowers2(int64(i-1)) + countPowers(remain) + (remain+1)*int64(i-1)
}

func getValue(n int64) int64 {
	var l int64
	r := n
	for l <= r {
		mid := (l + r) >> 1
		v := countOnes(mid)
		if v > n {
			r = mid - 1
		} else {
			if countOnes(mid+1) > n {
				l = mid
				break
			} else {
				l = mid + 1
				if r < l {
					r = l
				}
			}
		}
	}
	powers := countPowers(l)
	remain := n - countOnes(l)
	if remain > 0 {
		t := l + 1
		var i int64 = 0
		for remain > 0 {
			if t&1 == 1 {
				powers += i
				remain--
			}
			t >>= 1
			i++
		}
	}
	return powers
}
