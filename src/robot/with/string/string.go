package string

func robotWithString(s string) string {
	n := len(s)
	suffix := make([]byte, n)
	suffix[n-1] = 123
	for i := n - 2; i >= 0; i-- {
		if s[i+1] < suffix[i+1] {
			suffix[i] = s[i+1]
		} else {
			suffix[i] = suffix[i+1]
		}
	}
	str := make([]byte, 0, n)
	out := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		var top byte
		if len(str) != 0 {
			top = str[len(str)-1]
		} else {
			top = 124
		}
		if top <= s[i] {
			if top <= suffix[i] {
				out = append(out, top)
				str = str[:len(str)-1]
				i--
			} else {
				str = append(str, s[i])
			}
		} else {
			if s[i] <= suffix[i] {
				out = append(out, s[i])
			} else {
				str = append(str, s[i])
			}
		}
	}
	j := len(str) - 1
	for j >= 0 {
		out = append(out, str[j])
		j--
	}
	return string(out)
}
