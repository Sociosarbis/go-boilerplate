package rooms

const mask int = 1e9 + 7

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)

	times := make([]int, n)

	firstTime := make([]int, n)

	ret := 0

	i := 0

	for i+1 != n {
		ret = (ret + 1) % mask
		times[i]++
		if times[i] == 1 {
			firstTime[i] = ret
		}
		if times[i]%2 == 0 {
			i++
		} else {
			delta := firstTime[i] - firstTime[nextVisit[i]]
			if delta < 0 {
				delta += mask
			}
			ret = (ret + delta) % mask
		}
	}

	return ret
}
