package servers

func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) >> 1
		if target < arr[mid] {
			if mid > 0 {
				r = mid - 1
			} else {
				break
			}
		} else if target > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}

func busiestServers(k int, arrival []int, load []int) []int {
	ret := []int{}
	max := 0
	counters := make([]int, k)
	idleServers := make([]int, k)
	taskEndTimes := []int{}
	endTimeToServers := map[int][]int{}
	for i := 0; i < k; i += 1 {
		idleServers[i] = i
	}
	for i, t := range arrival {
		for j, endTime := range taskEndTimes {
			if endTime <= t {
				for _, i := range endTimeToServers[endTime] {
					l := binarySearch(idleServers, i)
					if l < len(idleServers) {
						idleServers = append(idleServers[:l+1], idleServers[l:]...)
						idleServers[l] = i
					} else {
						idleServers = append(idleServers, i)
					}
				}
				delete(endTimeToServers, endTime)
				if j == len(taskEndTimes) {
					taskEndTimes = taskEndTimes[:0]
				}
			} else {
				taskEndTimes = taskEndTimes[j:]
				break
			}
		}

		if len(idleServers) != 0 {
			l := binarySearch(idleServers, i%k) % len(idleServers)
			endTime := t + load[i]
			server := idleServers[l]
			counters[server] += 1
			if counters[server] == max {
				ret = append(ret, server)
			} else if counters[server] > max {
				ret = []int{server}
				max = counters[server]
			}
			if l+1 == len(idleServers) {
				idleServers = idleServers[:l]
			} else {
				idleServers = append(idleServers[:l], idleServers[l+1:]...)
			}
			if _, ok := endTimeToServers[endTime]; ok {
				endTimeToServers[endTime] = append(endTimeToServers[endTime], server)
			} else {
				endTimeToServers[endTime] = []int{server}
				l = binarySearch(taskEndTimes, endTime)
				if l < len(taskEndTimes) {
					taskEndTimes = append(taskEndTimes[:l+1], taskEndTimes[l:]...)
					taskEndTimes[l] = endTime
				} else {
					taskEndTimes = append(taskEndTimes, endTime)
				}
			}
		}
	}
	return ret
}
