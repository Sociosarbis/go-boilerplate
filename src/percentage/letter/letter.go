package letter

func percentageLetter(s string, letter byte) int {
	n := len(s)
	var c int
	for i := 0; i < n; i++ {
		if s[i] == letter {
			c++
		}
	}
	return c * 100 / n
}
