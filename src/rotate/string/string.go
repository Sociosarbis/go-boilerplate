package string

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	s = s + s
	for i := 1; i < len(goal); i++ {
		if s[i:i+len(goal)] == goal {
			return true
		}
	}
	return false
}
