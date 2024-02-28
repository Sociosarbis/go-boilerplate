package increments

func minIncrements(n int, cost []int) int {
	ret := 0
	for i := n - 1; i >= 2; i -= 2 {
		left := cost[i-1]
		right := cost[i]
		p := (i - 2) / 2
		if left > right {
			ret += left - right
			cost[p] += left
		} else if left < right {
			ret += right - left
			cost[p] += right
		} else {
			cost[p] += left
		}
	}
	return ret
}
