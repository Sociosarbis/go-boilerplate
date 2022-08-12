package people

func groupThePeople(groupSizes []int) [][]int {
	ret := make([][]int, len(groupSizes)+1)
	groupIdIndex := make([]int, len(groupSizes)+1)
	count := 0
	for i := 0; i < len(groupSizes); i++ {
		index := groupIdIndex[groupSizes[i]]
		if index == 0 || len(ret[index]) >= groupSizes[i] {
			index = count + 1
			groupIdIndex[groupSizes[i]] = index
			count++
		}
		ret[index] = append(ret[index], i)
	}
	return ret[1 : count+1]
}
