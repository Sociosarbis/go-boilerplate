package asterisks

func countAsterisks(s string) int {
	state := false
	ret := 0
	for _, c := range s {
		if c == '*' {
			if !state {
				ret++
			}
		} else if c == '|' {
			state = !state
		}
	}
	return ret
}
