package candies3

func distributeCandies(n int, limit int) int64 {
	i := n - 2*limit
	if i < 0 {
		i = 0
	}
	var ret int64
	end := limit
	if end > n {
		end = n
	}
	for ; i <= end; i++ {
		min := n - i - limit
		if min < 0 {
			min = 0
		}
		max := n - i
		if max > limit {
			max = limit
		}
		ret += int64(max - min + 1)
	}
	return ret
}
