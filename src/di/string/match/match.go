package match

func diStringMatch(s string) []int {
	l := 0
	r := len(s)
	j := 0
	ret := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			ret[j] = l
			l++
		} else {
			ret[j] = r
			r--
		}
		j++
	}
	ret[j] = l
	return ret
}
