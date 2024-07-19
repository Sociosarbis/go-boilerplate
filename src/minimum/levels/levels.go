package levels

func minimumLevels(possible []int) int {
	var b int
	for _, p := range possible {
		if p == 1 {
			b++
		} else {
			b--
		}
	}
	n := len(possible)
	var a int
	for i := 0; i+1 < n; i++ {
		if possible[i] == 1 {
			a++
			b--
		} else {
			a--
			b++
		}
		if a > b {
			return i + 1
		}
	}
	return -1
}
