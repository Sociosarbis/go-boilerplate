package awesome

func longestAwesome(s string) int {
	var mask int
	states := map[int]int{}
	for i, c := range s {
		bit := 1 << (c - 48)
		if mask&bit != 0 {
			mask -= bit
		} else {
			mask |= bit
		}
		states[mask] = i
	}

	mask = 0
	var ret int
	var odd_count int
	for i, c := range s {
		if j, ok := states[mask]; ok {
			if j-i+1 > ret {
				ret = j - i + 1
			}
		}
		for k := 0; k < 10; k++ {
			if mask&(1<<k) != 0 {
				if j, ok := states[mask-(1<<k)]; ok {
					if j-i+1 > ret {
						ret = j - i + 1
					}
				}
			}
		}
		for k := 0; k < 10; k++ {
			if mask&(1<<k) == 0 {
				if j, ok := states[mask|(1<<k)]; ok {
					if j-i+1 > ret {
						ret = j - i + 1
					}
				}
			}
		}
		bit := 1 << (c - 48)
		if mask&bit != 0 {
			mask -= bit
			odd_count--
		} else {
			mask |= bit
			odd_count++
		}
	}
	return ret
}
