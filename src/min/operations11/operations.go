package operations11

func getCounts(num int) []int {
	var mask int
	count := 0b11
	out := make([]int, 16)
	var i int
	for num >= mask+count {
		out[i] = count
		mask |= count
		count <<= 2
		i++
	}
	if num > mask {
		out[i] = num - mask
	}
	return out
}

func substractCounts(a []int, b []int) []int {
	for i := 0; i < 16; i++ {
		a[i] -= b[i]
	}
	return a
}

func count(counts []int) int64 {
	var i = 15
	var out int64
	for i >= 0 {
		if counts[i] == 0 {
			i--
			continue
		}
		if counts[i]%2 == 0 {
			out += int64(counts[i] / 2)
			if i > 0 {
				counts[i-1] += counts[i]
			}
			i--
		} else {
			if counts[i] != 1 {
				out += int64(counts[i] / 2)
				if i > 0 {
					counts[i-1] += counts[i] - 1
				}
				counts[i] = 1
			} else {
				j := i - 1
				for ; j >= 0; j-- {
					if counts[j] != 0 {
						break
					}
				}
				if j >= 0 {
					counts[i-1]++
					counts[j]--
					if j > 0 {
						counts[j-1]++
					}
					out++
				} else {
					out += int64(i + 1)
					break
				}
				i--
			}
		}
	}
	return out
}

func minOperations(queries [][]int) int64 {
	var out int64
	for _, query := range queries {
		b, a := getCounts(query[0]-1), getCounts(query[1])
		c := count(substractCounts(a, b))
		out += c
	}
	return out
}
