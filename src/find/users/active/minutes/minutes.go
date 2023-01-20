package minutes

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	idToMins := map[int]map[int]struct{}{}
	ret := make([]int, k)
	for _, log := range logs {
		if _, ok := idToMins[log[0]]; !ok {
			idToMins[log[0]] = map[int]struct{}{}
		}
		if _, ok := idToMins[log[0]][log[1]]; !ok {
			idToMins[log[0]][log[1]] = struct{}{}
		}
	}
	for _, v := range idToMins {
		if len(v)-1 < len(ret) {
			ret[len(v)-1]++
		}
	}
	return ret
}
