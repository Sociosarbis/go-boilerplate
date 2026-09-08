package commas

func countCommas(n int) int {
	var min int = 1e3
	if n < min {
		return 0
	}
	return n - min + 1
}
