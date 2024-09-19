package string

func longestContinuousSubstring(s string) int {
	ret := 1
	temp := 1
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1]+1 {
			temp++
		} else {
			if temp > ret {
				ret = temp
			}
			temp = 1
		}
	}
	if temp > ret {
		ret = temp
	}
	return ret
}
