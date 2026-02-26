package steps

func numSteps(s string) int {
	var out, next int
	n := len(s)
	for i := n - 1; i > 0; i-- {
		if s[i] == '0' {
			if next == 1 {
				out += 2
			} else {
				out++
			}
		} else {
			if next == 1 {
				out++
			} else {
				out += 2
				next = 1
			}
		}
	}
	if next == 1 {
		return out + 1
	}
	return out
}
