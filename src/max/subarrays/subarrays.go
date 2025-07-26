package subarrays

import (
	"sort"
)

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	for _, it := range conflictingPairs {
		if it[0] > it[1] {
			it[0], it[1] = it[1], it[0]
		}
	}
	sort.Slice(conflictingPairs, func(i, j int) bool {
		return conflictingPairs[i][0] < conflictingPairs[j][0] || (conflictingPairs[i][0] == conflictingPairs[j][0] && conflictingPairs[i][1] < conflictingPairs[j][1])
	})
	m := len(conflictingPairs)
	prevEnd := n + 1
	end := n + 1
	from := n
	j := m - 1
	var out int64
	var delta int64
	var temp int64
	for i := n; i >= 1; i-- {
		for j >= 0 && conflictingPairs[j][0] >= i {
			if conflictingPairs[j][1] <= end {
				if from > i {
					temp += int64(from-i) * int64(prevEnd-end)
					if temp > delta {
						delta = temp
					}
				}
				temp = 0
				from = i
				prevEnd, end = end, conflictingPairs[j][1]
			} else if conflictingPairs[j][1] < prevEnd {
				if from > i {
					temp += int64(from-i) * int64(prevEnd-end)
					if temp > delta {
						delta = temp
					}
				}
				prevEnd = conflictingPairs[j][1]
				from = i
			}
			j--
		}
		out += int64(end - i)
	}
	temp += int64(from) * int64(prevEnd-end)
	if temp > delta {
		delta = temp
	}
	return out + delta
}
