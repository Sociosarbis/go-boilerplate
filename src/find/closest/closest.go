package closest

func findClosest(words []string, word1 string, word2 string) int {
	i1 := -1
	i2 := -1
	ret := len(words)
	for i, word := range words {
		if word == word1 {
			i1 = i
		} else if word == word2 {
			i2 = i
		}

		if i1 != -1 && i2 != -1 {
			var temp int
			if i1 > i2 {
				temp = i1 - i2
			} else {
				temp = i2 - i1
			}
			if temp < ret {
				ret = temp
			}
		}
	}
	return ret
}
