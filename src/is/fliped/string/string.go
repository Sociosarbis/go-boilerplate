package string

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	repeat := s1 + s1
	for i := 1; i < len(s2); i++ {
		if repeat[i:i+len(s2)] == s2 {
			return true
		}
	}
	return false
}
