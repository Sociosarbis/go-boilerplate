package segments

func countSegments(s string) int {
	ret := 0
	count := 0
	for _, c := range s {
		if c != ' ' {
			count += 1
		} else {
			if count != 0 {
				ret += 1
				count = 0
			}
		}
	}
	if count != 0 {
		ret += 1
	}
	return ret
}
