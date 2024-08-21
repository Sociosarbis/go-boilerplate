package number

func count2(n int64, k int) int64 {
	i := 1
	step := (n + 1) >> 1
	var ret int64

	for n != 0 {
		if n&1 == 1 && i%k == 0 {
			ret += step
		}
		i++
		n >>= 1
	}
	return ret
}

func count(n int64, k int) int64 {
	if n == 0 {
		return 0
	}
	i := 0
	o := n
	for n != 0 {
		i++
		n >>= 1
	}
	var m int64 = 1 << (i - 1)
	ret := count2(m-1, k) + count(o-m, k)
	if i%k == 0 {
		ret += o - m + 1
	}
	return ret
}

func findMaximumNumber(k int64, x int) int64 {
	var l int64 = 1
	var r int64 = 1000000000000000
	for l <= r {
		mid := (l + r) >> 1
		v := count(mid, x)
		if v > k {
			r = mid - 1
		} else {
			if count(mid+1, x) > k {
				return mid
			} else {
				l = mid + 1
				if l > r {
					r = l
				}
			}
		}
	}
	return l
}
