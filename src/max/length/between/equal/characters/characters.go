package characters

func maxLengthBetweenEqualCharacters(s string) int {
	firstIndices := make([]int, 26)
	for i := range firstIndices {
		firstIndices[i] = -1
	}
	ret := -1
	for i, c := range s {
		index := byte(c - 97)
		if firstIndices[index] == -1 {
			firstIndices[index] = i
		} else {
			if i-firstIndices[index]-1 > ret {
				ret = i - firstIndices[index] - 1
			}
		}
	}
	return ret
}
