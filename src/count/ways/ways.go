package ways

import "sort"

const mask int = 1e9 + 7

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	n := len(ranges)
	groups := 1
	end := ranges[0][1]

	for i := 1; i < n; i++ {
		if ranges[i][0] > end {
			groups++
		}
		if ranges[i][1] > end {
			end = ranges[i][1]
		}
	}

	ret := 1

	for groups != 0 {
		ret = (ret * 2) % mask
		groups--
	}
	return ret
}
