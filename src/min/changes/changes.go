package changes

func minChanges(n int, k int) int {
	diff := n ^ k
	if diff&n != diff {
		return -1
	}
	var ret int
	for diff != 0 {
		diff &= diff - 1
		ret++
	}
	return ret
}
