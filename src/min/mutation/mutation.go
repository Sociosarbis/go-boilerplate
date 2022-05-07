package mutation

func isLessThanOneLetterDiff(s1 string, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count <= 1
}

func minMutation(start string, end string, bank []string) int {
	m := make(map[string]bool, len(bank))
	for _, s := range bank {
		m[s] = true
	}
	bfs := []string{start}
	ret := 0
	for len(bfs) != 0 {
		n := len(bfs)
		for i := 0; i < n; i++ {
			if bfs[i] == end {
				return ret
			}
			keysToAdd := []string{}
			for key := range m {
				if isLessThanOneLetterDiff(bfs[i], key) {
					keysToAdd = append(keysToAdd, key)
				}
			}
			for _, key := range keysToAdd {
				delete(m, key)
				bfs = append(bfs, key)
			}
		}
		ret++
		bfs = bfs[n:]
	}
	return -1
}
