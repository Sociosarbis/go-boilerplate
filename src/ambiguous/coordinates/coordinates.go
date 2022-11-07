package coordinates

func getNums(s string) []string {
	ret := []string{}
	hasPrefixZero := s[0] == '0'
	for i := 0; i < len(s); i++ {
		if i+1 != len(s) {
			if s[len(s)-1] != '0' {
				ret = append(ret, s[:i+1]+"."+s[i+1:])
			}
		} else {
			ret = append(ret, s[:i+1])
		}
		if hasPrefixZero {
			break
		}
	}
	return ret
}

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	ret := []string{}
	for i := 1; i < len(s); i++ {
		left := getNums(s[:i])
		if len(left) != 0 {
			right := getNums(s[i:])
			if len(right) != 0 {
				for j := 0; j < len(left); j++ {
					for k := 0; k < len(right); k++ {
						ret = append(ret, "("+left[j]+", "+right[k]+")")
					}
				}
			}
		}
	}
	return ret
}
