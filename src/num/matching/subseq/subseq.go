package subseq

func numMatchingSubseq(s string, words []string) int {
	m := make([][]int, 26)
	for i := 0; i < 26; i++ {
		m[i] = []int{}
	}
	for i, c := range s {
		m[c-97] = append(m[c-97], i)
	}
	ret := 0
loop:
	for _, word := range words {
		from := 0
		for _, c := range word {
			index := c - 97
			if len(m[index]) != 0 {
				l := 0
				r := len(m[index]) - 1
				for l <= r {
					mid := (l + r) / 2
					if m[index][mid] >= from {
						if mid > 0 && m[index][mid-1] >= from {
							r = mid - 1
						} else {
							l = mid
							break
						}
					} else {
						l = mid + 1
					}
				}
				if l >= len(m[index]) {
					continue loop
				}
				from = m[index][l] + 1
			} else {
				continue loop
			}
		}
		ret++
	}
	return ret
}
