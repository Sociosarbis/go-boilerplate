package characters

func rearrangeCharacters(s string, target string) int {
	m1 := make([]int, 26)
	for _, c := range s {
		m1[c-97]++
	}
	m2 := make(map[byte]int, 26)
	for i := 0; i < len(target); i++ {
		if c, ok := m2[target[i]]; ok {
			m2[target[i]] = c + 1
		} else {
			m2[target[i]] = 1
		}
	}
	ret := len(s) / len(target)

	for k, v := range m2 {
		temp := m1[k-97] / v
		if ret > temp {
			ret = temp
		}
	}
	return ret
}
