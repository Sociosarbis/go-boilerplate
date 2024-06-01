package candies

func distributeCandies(n int, limit int) int {
	var s1 int
	if limit > n {
		s1 = n
	} else {
		s1 = limit
	}
	var ret int
	for i := s1; i >= 0; i-- {
		remain := n - i
		var s2 int
		if limit > remain {
			s2 = remain
		} else {
			s2 = limit
		}
		var end int
		if remain >= limit {
			end = remain - limit
		} else {
			end = 0
		}
		if s2 < end {
			break
		} else {
			ret += s2 - end + 1
		}
	}
	return ret
}
