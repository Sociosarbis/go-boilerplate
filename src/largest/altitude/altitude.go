package altitude

func largestAltitude(gain []int) int {
	ret := 0
	temp := 0
	for _, num := range gain {
		temp += num
		if temp > ret {
			ret = temp
		}
	}
	return ret
}
