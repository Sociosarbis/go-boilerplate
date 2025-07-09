package time

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	var prev, out int
	n := len(startTime)
	prefix := make([]int, n+1)
	for i, start := range startTime {
		space := start - prev
		if i > 0 {
			from := i - k
			if from < 0 {
				from = 0
			}
			temp := prefix[i] - prefix[from] + space
			if temp > out {
				out = temp
			}
		} else {
			if space > out {
				out = space
			}
		}
		prefix[i+1] = prefix[i] + space
		prev = endTime[i]
	}
	space := eventTime - prev
	from := n - k
	temp := prefix[n] - prefix[from] + space
	if temp > out {
		out = temp
	}
	return out
}
