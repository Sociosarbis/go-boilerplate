package equal

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i1 := 0
	j1 := 0
	i2 := 0
	j2 := 0
	for i1 < len(word1) {
		if j1 < len(word1[i1]) {
			c1 := word1[i1][j1]
			if i2 < len(word2) {
				if j2 < len(word2[i2]) {
					c2 := word2[i2][j2]
					if c1 != c2 {
						return false
					}
					j1++
					j2++
				} else {
					i2++
					j2 = 0
				}
			} else {
				return false
			}
		} else {
			j1 = 0
			i1++
		}
	}

	if i2 < len(word2) {
		if j2 >= len(word2[i2]) {
			i2++
		}
		if i2 < len(word2) {
			return false
		}
	}
	return true
}
