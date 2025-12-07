package odds

func countOdds(low int, high int) int {
	count := high - low + 1
	if count&1 == 1 && low&1 == 1 {
		return count/2 + 1
	}
	return count / 2
}
