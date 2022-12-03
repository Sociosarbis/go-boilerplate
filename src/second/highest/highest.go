package highest

func secondHighest(s string) int {
	ret := []int{-1, -1}
	for _, c := range s {
		v := int(c - 48)
		if v >= 0 && v < 10 && v != ret[0] && v != ret[1] {
			i := 2
			for ; i > 0; i-- {
				if v < ret[i-1] {
					break
				}
			}
			if i < 2 {
				copy(ret[i+1:], ret[i:])
				ret[i] = v
			}
		}
	}
	return ret[1]
}
