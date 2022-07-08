package chips

func minCostToMoveChips(position []int) int {
	evenMinusOdd := 0
	for _, p := range position {
		if p&1 == 1 {
			evenMinusOdd--
		} else {
			evenMinusOdd++
		}
	}

	if evenMinusOdd > 0 {
		return (len(position) - evenMinusOdd) / 2
	}
	return (len(position) + evenMinusOdd) / 2
}
