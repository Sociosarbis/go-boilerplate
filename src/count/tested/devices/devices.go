package devices

func countTestedDevices(batteryPercentages []int) int {
	var ret int
	for _, p := range batteryPercentages {
		if p > ret {
			ret++
		}
	}
	return ret
}
