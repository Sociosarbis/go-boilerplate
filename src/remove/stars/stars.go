package stars

func removeStars(s string) string {
	n := len(s)
	ret := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			if len(ret) > 0 {
				ret = ret[:len(ret)-1]
			}
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}
