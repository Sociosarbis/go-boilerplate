package moves

func minimumMoves(s string) int {
	i := 0
	ret := 0
	for i < len(s) {
		if s[i] == 'X' {
			i += 3
			ret++
		} else {
			i++
		}
	}
	return ret
}
