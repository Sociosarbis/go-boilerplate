package string

func scoreOfString(s string) int {
	var ret int
	n := len(s)
	for i := 1; i < n; i++ {
		v := int(s[i]) - int(s[i-1])
		if v >= 0 {
			ret += v
		} else {
			ret -= v
		}
	}
	return ret
}
