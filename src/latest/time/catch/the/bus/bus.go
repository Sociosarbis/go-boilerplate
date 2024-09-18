package bus

import (
	"sort"
)

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	var ret int

	var remainSpace int

	var j int

	n := len(passengers)

	for _, bus := range buses {
		remainSpace = capacity
		for remainSpace > 0 && (j < n && passengers[j] <= bus) {
			if j == 0 || passengers[j-1]+1 != passengers[j] {
				ret = passengers[j] - 1
			}
			j++
			remainSpace--
		}
		if remainSpace > 0 && (j == 0 || passengers[j-1] != bus) {
			ret = bus
		}
	}
	return ret
}
