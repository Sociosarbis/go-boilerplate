package str2

func processStr(s string, k int64) byte {
	iToN := make([]int64, len(s))
	var n int64
	for i, c := range s {
		if c == '*' {
			if n > 0 {
				n--
			}
		} else if c == '#' {
			n *= 2
		} else if c != '%' {
			n++
		}
		iToN[i] = n
	}
	for i := len(s) - 1; i >= 0; i-- {
		if k+1 > iToN[i] {
			return '.'
		}
		if s[i] == '#' {
			if k*2 >= iToN[i] {
				k -= iToN[i] / 2
			}
		} else if s[i] == '%' {
			k = iToN[i] - 1 - k
		} else if s[i] != '*' {
			if k+1 == iToN[i] {
				return s[i]
			}
		}
	}
	return '*'
}
