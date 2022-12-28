package length

func minimumLength(s string) int {
	if len(s) == 1 {
		return 1
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[j] != s[i] {
			break
		}
		for i < j && s[i+1] == s[i] {
			i++
		}
		if i == j {
			return 0
		}
		for j > 0 && s[j-1] == s[j] {
			j--
		}
		i++
		j--
	}
	return j - i + 1
}
