package digits

func clearDigits(s string) string {
	n := len(s)
	ret := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(ret) > 0 {
				ret = ret[:len(ret)-1]
			}
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}
