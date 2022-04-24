package gap

func binaryGap(n int) int {
	i := 0
	j := 0
	for n != 0 && n&1 != 1 {
		i++
		n >>= 1
	}

	ret := 0

	if n != 0 {
		j = i
		for n != 0 {
			if n&1 == 1 {
				if i-j > ret {
					ret = i - j
				}
				j = i
			}
			n >>= 1
			i++
		}
	}
	return ret
}
