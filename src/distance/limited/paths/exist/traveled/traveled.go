package traveled

func distanceTraveled(mainTank int, additionalTank int) int {
	temp := 0
	for mainTank >= 5 {
		times := mainTank / 5
		temp += times
		if times > additionalTank {
			times = additionalTank
		}
		additionalTank -= times
		mainTank = mainTank%5 + times
	}
	return ((temp * 5) + mainTank) * 10
}
