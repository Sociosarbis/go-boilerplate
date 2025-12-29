package transition

func dfs(bottomToTops map[string][]byte, current []byte, i int, next *[]byte) bool {
	n := len(current)
	if n == 1 {
		return true
	}
	if i+1 >= n {
		return dfs(bottomToTops, *next, 0, &[]byte{})
	}
	prefix := string([]byte{current[i], current[i+1]})
	n = len(*next)
	for _, top := range bottomToTops[prefix] {
		*next = append(*next, top)
		temp := dfs(bottomToTops, current, i+1, next)
		if temp {
			return temp
		}
		*next = (*next)[:n]
	}
	return false
}

func pyramidTransition(bottom string, allowed []string) bool {
	bottomToTops := map[string][]byte{}
	for _, it := range allowed {
		prefix := it[:2]
		if tops, ok := bottomToTops[prefix]; ok {
			bottomToTops[prefix] = append(tops, it[2])
		} else {
			bottomToTops[prefix] = []byte{it[2]}
		}
	}
	return dfs(bottomToTops, []byte(bottom), 0, &[]byte{})
}
