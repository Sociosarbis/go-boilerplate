package target

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var ret int
	for _, hour := range hours {
		if hour >= target {
			ret++
		}
	}
	return ret
}
