package size2

func minDeletionSize(strs []string) int {
	var out int
	var start int
	m := len(strs)
	lessThanNext := make([]bool, m)
	n := len(strs[0])
	end := m - 1
loop:
	for j := 0; j < n; j++ {
		for i := start; i < end; i++ {
			if lessThanNext[i] || strs[i][j] <= strs[i+1][j] {
				continue
			} else {
				out++
				continue loop
			}
		}
		for i := start; i < end; i++ {
			if strs[i][j] < strs[i+1][j] {
				lessThanNext[i] = true
			}
		}
		for i := start; i < end; i++ {
			if lessThanNext[i] {
				start++
			} else {
				break
			}
		}
	}
	return out
}
