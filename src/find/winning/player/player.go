package player

func findWinningPlayer(skills []int, k int) int {
	var ret int
	n := len(skills)
	if skills[0] < skills[1] {
		ret = 1
	} else {
		ret = 0
	}
	acc := 1
	if acc == k {
		return ret
	}
	for j := 2; j < n; j++ {
		if skills[ret] < skills[j] {
			ret = j
			acc = 1
		} else {
			acc++
			if acc == k {
				return ret
			}
		}
	}
	return ret
}
