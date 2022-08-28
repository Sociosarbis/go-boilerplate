package fzf

func getZeroes(num int64) int {
	base := 1
	ret := 0
	for num >= 5 {
		nextNum := num / 5
		if nextNum >= 5 {
			ret += base * int(nextNum-nextNum/5)
		} else {
			ret += base * int(nextNum)
		}
		base++
		num = nextNum
	}
	return ret
}

func preimageSizeFZF(k int) int {
	var l int64 = 0
	var r int64 = 4000000015
	for l <= r {
		mid := (l + r) / 2
		zeroes := getZeroes(mid)
		if zeroes < k {
			l = mid + 1
		} else if zeroes > k {
			r = mid - 1
		} else {
			return 5
		}
	}
	return 0
}
