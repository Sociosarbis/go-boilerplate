package valid

func minAddToMakeValid(s string) int {
	ret := 0
	temp := 0
	for _, c := range s {
		if c == '(' {
			temp++
		} else {
			temp--
		}
		if temp < 0 {
			ret -= temp
			temp = 0
		}
	}
	ret += temp
	return ret
}
