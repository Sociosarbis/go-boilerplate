package valid2

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	var s int
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		if c >= 'a' && c <= 'z' {
			if s&1 == 0 {
				if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
					s |= 1
				}
			}
			if s&2 == 0 {
				if !(c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
					s |= 2
				}
			}
			continue
		}
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return s == 3
}
