package variance

const fullMask int = (1 << 26) - 1

func largestVariance(s string) int {
	var mask int
	for _, c := range s {
		mask |= 1 << (c - 97)
		if mask == fullMask {
			break
		}
	}
	n := len(s)
	var ret int
	for i := 0; i < 25; i++ {
		if mask&(1<<i) != 0 {
			a := byte(i + 97)
			for j := i + 1; j < 26; j++ {
				if mask&(1<<j) != 0 {
					b := byte(j + 97)
					for c := 0; c < 2; c++ {
						var c1 int
						var c2 int
						for k := 0; k < n; k++ {
							if s[k] == a {
								c1++
								if c2 > 0 {
									if c1-c2 > ret {
										ret = c1 - c2
									}
								} else {
									if c1-1 > ret {
										ret = c1 - 1
									}
								}
							} else if s[k] == b {
								if c2 == c1 {
									if c1 != 0 {
										c1 = 0
										c2 = 0
									}
								} else {
									c2++
								}
							}
						}
						a, b = b, a
					}
				}
			}
		}
	}
	return ret
}
