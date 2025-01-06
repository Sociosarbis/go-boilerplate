package consecutive

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	ret := special[0] - bottom
	n := len(special)
	for i := 1; i < n; i++ {
		v := special[i] - special[i-1] - 1
		if v > ret {
			ret = v
		}
	}
	v := top - special[n-1]
	if v > ret {
		return v
	}
	return ret
}
