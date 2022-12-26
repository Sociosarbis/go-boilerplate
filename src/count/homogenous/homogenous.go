package homogenous

func countHomogenous(s string) int {
	var ret int64
	var temp int64 = 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			temp++
		} else {
			ret += (1 + temp) * temp / 2
			if ret > 1e9+7 {
				ret %= 1e9 + 7
			}
			temp = 1
		}
	}
	ret += (1 + temp) * temp / 2
	if ret > 1e9+7 {
		ret %= 1e9 + 7
	}
	return int(ret)
}
