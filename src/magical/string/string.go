package string

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}
	s1 := make([]byte, n)
	s2 := make([]byte, n)
	copy(s1[:3], "122")
	copy(s2[:2], "12")
	i := 3
	j := 2
	ret := 1
	for i < n {
		s2[j] = s1[j]
		var c byte
		if s1[i-1] == '2' {
			c = '1'
		} else {
			c = '2'
		}
		s1[i] = c
		i++
		if c == '1' {
			ret++
		}
		if s2[j] == '2' {
			if i < n {
				s1[i] = c
				i++
				if c == '1' {
					ret++
				}
			}
		}
		j++
	}
	return ret
}
