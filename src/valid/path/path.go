package path

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	groupId := 0
	groups := map[int][]int{}
	groupIds := make([]int, n)
	for _, e := range edges {
		id1 := groupIds[e[0]]
		id2 := groupIds[e[1]]
		if id1 == 0 {
			if id2 == 0 {
				groupId++
				groupIds[e[0]] = groupId
				groupIds[e[1]] = groupId
				groups[groupId] = []int{e[0], e[1]}
			} else {
				groupIds[e[0]] = id2
				groups[id2] = append(groups[id2], e[0])
			}
		} else {
			if id2 == 0 {
				groupIds[e[1]] = id1
				groups[id1] = append(groups[id1], e[1])
			} else {
				if id1 != id2 {
					for _, v := range groups[id2] {
						groupIds[v] = id1
						groups[id1] = append(groups[id1], v)
					}
					delete(groups, id2)
				}
			}
		}
	}
	if groupIds[source] != 0 && groupIds[destination] != 0 {
		return groupIds[source] == groupIds[destination]
	}
	return false
}
