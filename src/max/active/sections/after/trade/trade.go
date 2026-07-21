package trade

func maxActiveSectionsAfterTrade(s string) int {
	var out int
	for _, c := range s {
		if c == '1' {
			out++
		}
	}
	counter := [2]int{}
	var temp int
	var st int
	for _, c := range s {
		if c == '1' {
			if st == 0 {
				if counter[0] != 0 {
					st = 1
				}
			} else if st == 2 {
				if counter[0]+counter[1] > temp {
					temp = counter[0] + counter[1]
				}
				counter[0], counter[1] = counter[1], 0
				st = 1
			}
		} else {
			if st == 0 {
				counter[0]++
			} else if st == 1 {
				counter[1] = 1
				st = 2
			} else {
				counter[1]++
			}
		}
	}
	if st == 2 {
		if counter[0]+counter[1] > temp {
			temp = counter[0] + counter[1]
		}
	}
	return out + temp
}
