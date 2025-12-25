package sum

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	var out int64
	for i, value := range happiness {
		if i < k && value > i {
			out += int64(value)
		} else {
			return out - int64(i-1)*int64(i)/2
		}
	}
	n := len(happiness)
	return out - int64(n-1)*int64(n)/2
}
