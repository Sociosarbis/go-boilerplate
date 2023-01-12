package evaluate2

func evaluate(s string, knowledge [][]string) string {
	if len(s) < 3 {
		return s
	}
	i := 0
	j := i
	m := map[string]string{}
	for _, p := range knowledge {
		m[p[0]] = p[1]
	}
	ret := ""
	for j < len(s) {
		if s[j] != '(' {
			j++
		} else {
			ret += s[i:j]
			j++
			i = j
			for j < len(s) {
				if s[j] != ')' {
					j++
				} else {
					break
				}
			}
			k := s[i:j]
			if w, ok := m[k]; ok {
				ret += w
			} else {
				ret += "?"
			}
			j++
			i = j
		}
	}
	if i != j {
		ret += s[i:j]
	}
	return ret
}
