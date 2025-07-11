package days

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	var prev, out int
	for _, meeting := range meetings {
		if meeting[0] > prev+1 {
			out += meeting[0] - prev - 1
		}
		if meeting[1] > prev {
			prev = meeting[1]
		}
	}
	if days > prev {
		return out + days - prev
	}
	return out
}
