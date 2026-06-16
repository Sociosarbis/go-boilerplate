package str

func processStr(s string) string {
	out := make([]byte, 0, len(s))
	for _, c := range s {
		if c == '*' {
			if len(out) > 0 {
				out = out[:len(out)-1]
			}
		} else if c == '#' {
			out = append(out, out...)
		} else if c == '%' {
			var l int
			r := len(out) - 1
			for l <= r {
				out[l], out[r] = out[r], out[l]
				l++
				r--
			}
		} else {
			out = append(out, byte(c))
		}
	}
	return string(out)
}
