package grid

func createGrid(m int, n int) []string {
	lastRow := make([]byte, n)
	for i := 0; i < n; i++ {
		lastRow[i] = '.'
	}
	out := make([]string, m)
	out[m-1] = string(lastRow)
	if m > 1 {
		sharedRow := make([]byte, n)
		sharedRow[0] = '.'
		for i := 1; i < n; i++ {
			sharedRow[i] = '#'
		}
		ss := string(sharedRow)
		for i := m - 2; i >= 0; i-- {
			out[i] = ss
		}
	}
	return out
}
