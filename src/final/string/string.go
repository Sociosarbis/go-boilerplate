package string

func reverse(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func finalString(s string) string {
	n := len(s)
	ret := make([]byte, 0, n)
	r := false
	for i := 0; i < n; i++ {
		if s[i] == 'i' {
			r = !r
		} else {
			if r {
				reverse(ret)
				r = false
			}
			ret = append(ret, s[i])
		}
	}
	if r {
		reverse(ret)
	}
	return string(ret)
}
