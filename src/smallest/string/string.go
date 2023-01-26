package string

func getSmallestString(n int, k int) string {
	x := 0
	r := n
	for x <= r {
		mid := (x + r) / 2
		if mid+(n-mid)*26 < k {
			r = mid - 1
		} else {
			if (mid + 1 + (n-mid-1)*26) >= k {
				x = mid + 1
				continue
			}
			x = mid
			break
		}
	}
	ret := make([]byte, n)
	i := 0
	for ; i < x; i++ {
		ret[i] = 'a'
	}
	if i < n {
		if x+(n-x)*26 > k {
			ret[i] = byte(122 - ((x + (n-x)*26) - k))
			i++
		}
	}
	for ; i < n; i++ {
		ret[i] = 'z'
	}
	return string(ret)
}
