package tax

func calculateTax(brackets [][]int, income int) float64 {
	var ret float64 = 0
	for i, bracket := range brackets {
		var val int
		if income <= bracket[0] {
			val = income
		} else {
			val = bracket[0]
		}
		if i == 0 {
			ret += float64(val) * float64(bracket[1]) / 100
		} else {
			ret += float64(val-brackets[i-1][0]) * float64(bracket[1]) / 100
		}
		if val == income {
			break
		}
	}
	return ret
}
