package segment

func checkOnesSegment(s string) bool {
	consumed := false
	for i := 0; i < len(s); i++ {
		if consumed {
			if s[i] == '1' {
				return false
			}
		} else {
			if s[i] != '1' {
				consumed = true
			}
		}
	}
	return true
}
