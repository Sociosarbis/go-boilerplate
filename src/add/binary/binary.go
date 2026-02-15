package binary

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(a, b byte) byte {
	return a + b
}

func reversed(s []byte) []byte {
	var l int
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}

func addBinary(a string, b string) string {
	n1 := len(a)
	n2 := len(b)
	n := max(n1, n2)
	out := make([]byte, 0, n+1)
	i := 1
	var remain byte = '0'
	for i <= n1 || i <= n2 {
		var sum byte
		if i <= n1 {
			if i <= n2 {
				sum = add(add(a[n1-i], b[n2-i]), remain)
			} else {
				sum = add(a[n1-i], remain)
			}
		} else {
			sum = add(b[n2-i], remain)
		}
		sum -= 96
		if sum&1 == 1 {
			out = append(out, '1')
		} else {
			out = append(out, '0')
		}
		if sum&2 != 0 {
			remain = '1'
		} else {
			remain = '0'
		}
		i++
	}
	if remain != '0' {
		out = append(out, remain)
	}
	return string(reversed(out))
}
