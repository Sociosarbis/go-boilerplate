package reformat

func reformat(s string) string {
	c1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			c1++
		}
	}
	c2 := len(s) - c1
	if c2 != c1 && c2+1 != c1 && c1+1 != c2 {
		return ""
	}
	var i1 int
	var i2 int
	if c1 >= c2 {
		i1 = 0
		i2 = 1
	} else {
		i1 = 1
		i2 = 0
	}
	ret := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			ret[i2] = s[i]
			i2 += 2
		} else {
			ret[i1] = s[i]
			i1 += 2
		}
	}
	return string(ret)
}
