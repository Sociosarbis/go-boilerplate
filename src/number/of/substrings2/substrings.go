package substrings2

func getNextChar(s string, from int, c byte) int {
	n := len(s)
	for i := from; i < n; i++ {
		if s[i] == c {
			return i
		}
	}
	return n
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func numberOfSubstrings(s string) int {
	n := len(s)
	a, b, c := getNextChar(s, 0, 'a'), getNextChar(s, 0, 'b'), getNextChar(s, 0, 'c')
	end := max(a, b, c)
	if end == n {
		return 0
	}
	out := n - end
	for i := 1; i < n; i++ {
		if i-1 == a {
			a = getNextChar(s, i, 'a')
		} else if i-1 == b {
			b = getNextChar(s, i, 'b')
		} else {
			c = getNextChar(s, i, 'c')
		}
		end = max(a, b, c)
		if end == n {
			break
		}
		out += n - end
	}
	return out
}
