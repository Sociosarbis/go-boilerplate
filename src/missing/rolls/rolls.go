package rolls

func missingRolls(rolls []int, mean int, n int) []int {
	total := mean * (len(rolls) + n)
	ret := make([]int, n)
	for _, r := range rolls {
		total -= r
	}

	if total < n || total > 6*n {
		return ret[:0]
	}
	avg := total / n
	rest := total % n
	for i := 0; i < n; i++ {
		if rest != 0 {
			ret[i] = avg + 1
			rest -= 1
		} else {
			ret[i] = avg
		}
	}
	return ret
}
