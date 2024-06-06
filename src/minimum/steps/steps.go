package steps

func minimumSteps(s string) int64 {
	var ret int64
	var index int
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			ret += int64(i - index)
			index++
		}
	}
	return ret
}
