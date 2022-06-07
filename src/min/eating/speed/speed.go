package speed

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := piles[0]
	for i := 1; i < len(piles); i++ {
		if r < piles[i] {
			r = piles[i]
		}
	}

	ret := r

	for l <= r {
		mid := (l + r) >> 1
		totalH := 0
		for i := 0; i < len(piles); i++ {
			if piles[i]%mid == 0 {
				totalH += piles[i] / mid
			} else {
				totalH += piles[i]/mid + 1
			}
		}
		if totalH <= h {
			if mid < ret {
				ret = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ret
}
