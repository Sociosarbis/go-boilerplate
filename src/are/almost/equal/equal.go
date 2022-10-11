package equal

func areAlmostEqual(s1 string, s2 string) bool {
	index := 0
	diff := make([]int, 2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if index == 2 {
				return false
			}
			diff[index] = i
			index++
		}
	}
	if index == 1 {
		return false
	}
	if index == 2 {
		return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
	}
	return true
}
