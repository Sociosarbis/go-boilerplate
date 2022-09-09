package operations

func minOperations(logs []string) int {
	ret := 0
	for _, log := range logs {
		if log[0] == '.' {
			if log[1] == '.' && len(log) == 3 {
				if ret > 0 {
					ret--
				}
				continue
			} else if len(log) == 2 {
				continue
			}
		}
		ret++
	}
	return ret
}
