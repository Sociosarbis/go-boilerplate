package names

const cn = 26

const fullMask = (1 << cn) - 1

func getByteMask(c byte) int {
	return 1 << (c - 'a')
}

func hasBit(mask int, i int) bool {
	return mask&(1<<i) != 0
}

func distinctNames(ideas []string) int64 {
	n := len(ideas)
	dict := make(map[string]int, n)

	for _, idea := range ideas {
		first := idea[0]
		suffix := idea[1:]
		if mask, ok := dict[suffix]; ok {
			dict[suffix] = mask | getByteMask(first)
		} else {
			dict[suffix] = getByteMask(first)
		}
	}

	var ret int64

	counter := make([][]int, cn)

	for i := 0; i < cn; i++ {
		counter[i] = make([]int, cn)
	}

	for _, mask := range dict {
		for i := 0; i < cn; i++ {
			if hasBit(mask, i) {
				for j := 0; j < cn; j++ {
					if !hasBit(mask, j) {
						counter[i][j]++
					}
				}
			}
		}
	}

	for i := 0; i < cn; i++ {
		for j := i + 1; j < cn; j++ {
			ret += int64(counter[i][j]) * int64(counter[j][i]) * 2
		}
	}

	return ret
}
