package operations

func minOperations(s string) int {
	ret := 0
	for i, c := range s {
		if i&1 == 0 {
			if c == '1' {
				ret++
			}
		} else {
			if c == '0' {
				ret++
			}
		}
	}
	temp := 0
	for i, c := range s {
		if i&1 == 0 {
			if c == '0' {
				temp++
			}
		} else {
			if c == '1' {
				temp++
			}
		}
	}
	if temp < ret {
		return temp
	}
	return ret
}
