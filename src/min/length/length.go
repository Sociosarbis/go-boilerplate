package length

func minLength(s string) int {
	i := 0
	for i+1 < len(s) {
		temp := s[i : i+2]
		if temp == "AB" || temp == "CD" {
			s = s[:i] + s[i+2:]
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return len(s)
}
