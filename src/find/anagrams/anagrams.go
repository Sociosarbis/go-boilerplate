package anagrams

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	symbol := [26]int{}

	mask := 0

	for _, c := range p {
		index := c - 97
		symbol[index] += 1
		mask |= 1 << index
	}
	for i := 0; i < len(p)-1; i += 1 {
		index := s[i] - 97
		symbol[index] -= 1
		if symbol[index] != 0 {
			mask |= 1 << index
		} else {
			mask -= 1 << index
		}
	}

	ret := make([]int, 0)

	for i := len(p) - 1; i < len(s); i += 1 {
		index := s[i] - 97
		symbol[index] -= 1
		if symbol[index] != 0 {
			mask |= 1 << index
		} else {
			mask -= 1 << index
		}
		popI := i + 1 - len(p)
		index = s[popI] - 97
		if mask == 0 {
			ret = append(ret, popI)
		}
		symbol[index] += 1
		if symbol[index] != 0 {
			mask |= 1 << index
		} else {
			mask -= 1 << index
		}
	}
	return ret
}
