package size

func minDeletionSize(strs []string) int {
	ret := 0
	m := len(strs)
	n := len(strs[0])
	for i := 0; i < n; i++ {
		prev := strs[0][i]
		for j := 1; j < m; j++ {
			if strs[j][i] < prev {
				ret++
				break
			}
			prev = strs[j][i]
		}
	}
	return ret
}
