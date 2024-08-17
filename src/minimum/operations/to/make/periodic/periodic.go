package periodic

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word) / k
	counter := make(map[string]int, n)
	m := len(word)
	max := 1
	for i := 0; i < m; i += k {
		if c, ok := counter[word[i:i+k]]; ok {
			counter[word[i:i+k]] = c + 1
			if c >= max {
				max = c + 1
			}
		} else {
			counter[word[i:i+k]] = 1
		}
	}
	return n - max
}
