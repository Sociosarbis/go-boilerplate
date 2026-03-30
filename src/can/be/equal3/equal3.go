package equal3

func checkStrings(s1 string, s2 string) bool {
	n := len(s1)
	counter := make(map[byte]int, 26)
	for i := 0; i < n; i += 2 {
		counter[s1[i]]++
		counter[s2[i]]--
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	for i := 1; i < n; i += 2 {
		counter[s1[i]]++
		counter[s2[i]]--
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true
}
