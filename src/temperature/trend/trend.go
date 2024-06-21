package trend

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	var ret int
	n := len(temperatureA)
	temp := 0
	for i := 1; i < n; i++ {
		stateA := getState(temperatureA[i], temperatureA[i-1])
		stateB := getState(temperatureB[i], temperatureB[i-1])
		if stateA == stateB {
			temp++
			if temp > ret {
				ret = temp
			}
		} else {
			temp = 0
		}
	}
	return ret
}

func getState(a int, b int) uint8 {
	if a > b {
		return 1
	} else if a < b {
		return 2
	}
	return 0
}
