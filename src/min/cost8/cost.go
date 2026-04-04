package cost8

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	var step int
	if startPos[0] < homePos[0] {
		step = 1
	} else if startPos[0] > homePos[0] {
		step = -1
	}
	var out int
	if step != 0 {
		for i := startPos[0] + step; ; i += step {
			out += rowCosts[i]
			if i == homePos[0] {
				break
			}
		}
		step = 0
	}
	if startPos[1] < homePos[1] {
		step = 1
	} else if startPos[1] > homePos[1] {
		step = -1
	}
	if step != 0 {
		for i := startPos[1] + step; ; i += step {
			out += colCosts[i]
			if i == homePos[1] {
				break
			}
		}
	}
	return out
}
