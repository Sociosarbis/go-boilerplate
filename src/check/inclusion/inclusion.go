package inclusion

func checkInclusion(s1 string, s2 string) bool {
	sign := [26]int{}

	for _, c := range s1 {
		sign[c-97]++
	}

	tmp := [26]int{}

	j := 0

	for i, c := range s2 {
		index := c - 97
		tmp[index]++
		for tmp[index] > sign[index] {
			index1 := s2[j] - 97
			tmp[index1]--
			j++
		}
		if i+1-j == len(s1) {
			return true
		}
	}
	return false
}
