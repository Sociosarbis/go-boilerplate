package string

func getSmallestString(s string, k int) string {
	ret := []byte(s)

	for i, c := range ret {
		d1 := c - 'a'
		d2 := 'z' - c + 1
		var d int
		if d1 < d2 {
			d = int(d1)
		} else {
			d = int(d2)
		}
		if d <= k {
			ret[i] = 'a'
			k -= d
		} else {
			ret[i] -= byte(k)
			k = 0
			break
		}
	}
	return string(ret)
}
