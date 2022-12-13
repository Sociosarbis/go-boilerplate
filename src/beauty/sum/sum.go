package sum

func beautySum(s string) int {
	ret := 0
	for i, c := range s {
		counter := make([]int, 26)
		counter[c-97]++
		min := 1
		max := 1
		for j := i + 1; j < len(s); j++ {
			index := s[j] - 97
			oldCount := counter[index]
			counter[index]++
			if counter[index] > max {
				max = counter[index]
			}
			if counter[index] == 1 {
				min = 1
			} else if oldCount == min {
				min = counter[index]
				for _, v := range counter {
					if v > 0 && v < min {
						min = v
						break
					}
				}
			}
			if max != min {
				ret += max - min
			}
		}
	}
	return ret
}
