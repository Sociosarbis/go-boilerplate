package winner

func isWinner(player1 []int, player2 []int) int {
	temp := 0
	mask := 0
	for _, num := range player1 {
		if mask&3 != 0 {
			temp += num * 2
		} else {
			temp += num
		}
		mask <<= 1
		if num == 10 {
			mask |= 1
		}
		mask &= 3
	}
	mask = 0
	for _, num := range player2 {
		if mask&3 != 0 {
			temp -= num * 2
		} else {
			temp -= num
		}
		mask <<= 1
		if num == 10 {
			mask |= 1
		}
		mask &= 3
	}
	if temp > 0 {
		return 1
	} else if temp < 0 {
		return 2
	}

	return 0
}
