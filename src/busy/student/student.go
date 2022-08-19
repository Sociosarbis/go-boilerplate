package student

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ret := 0
	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ret++
		}
	}
	return ret
}
