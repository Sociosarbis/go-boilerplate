package time2

import (
	"fmt"
	"sort"
)

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	maxSpaces := make([][2]int, 0, n+1)
	var prevSpace int
	for i, start := range startTime {
		if start > prevSpace {
			maxSpaces = append(maxSpaces, [2]int{start - prevSpace, i})
		}
		prevSpace = endTime[i]
	}
	if eventTime > prevSpace {
		maxSpaces = append(maxSpaces, [2]int{eventTime - prevSpace, n})
	}
	sort.Slice(maxSpaces, func(i, j int) bool {
		return maxSpaces[i][0] > maxSpaces[j][0]
	})
	if len(maxSpaces) > 3 {
		maxSpaces = maxSpaces[:3]
	}
	fmt.Println(maxSpaces)
	var prev, out int

	for i, start := range startTime {
		space := start - prev
		if i > 0 {
			var temp int
			for j := 0; j < len(maxSpaces); j++ {
				if endTime[i-1]-startTime[i-1] <= maxSpaces[j][0] && maxSpaces[j][1] != i-1 && maxSpaces[j][1] != i {
					if i > 1 {
						temp = startTime[i] - endTime[i-2]
					} else {
						temp = startTime[i]
					}
					if temp > out {
						out = temp
					}
				}
			}
			temp = prevSpace + space
			if temp > out {
				out = temp
			}
		} else {
			if space > out {
				out = space
			}
		}
		prev, prevSpace = endTime[i], space
	}
	space := eventTime - prev
	var temp int
	if n > 1 {
		for j := 0; j < len(maxSpaces); j++ {
			if endTime[n-1]-startTime[n-1] <= maxSpaces[j][0] && maxSpaces[j][1] != n-1 && maxSpaces[j][1] != n {
				temp = eventTime - endTime[n-2]
				if temp > out {
					return temp
				}
			}
		}
	}
	temp = prevSpace + space
	if temp > out {
		out = temp
	}
	return out
}
