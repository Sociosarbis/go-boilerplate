package distance

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	indexToGroup := make([]int, n)
	groups := [][]int{}
	for _, swap := range allowedSwaps {
		i, j := swap[0], swap[1]
		if indexToGroup[i] == 0 {
			if indexToGroup[j] == 0 {
				id := len(groups) + 1
				groups = append(groups, []int{i, j})
				indexToGroup[i], indexToGroup[j] = id, id
			} else {
				id := indexToGroup[j]
				indexToGroup[i] = id
				groups[id-1] = append(groups[id-1], i)
			}
		} else {
			id := indexToGroup[i]
			if indexToGroup[j] == 0 {
				indexToGroup[j] = id
				groups[id-1] = append(groups[id-1], j)
			} else if indexToGroup[j] != id {
				oldId := indexToGroup[j]
				for _, index := range groups[oldId-1] {
					groups[id-1] = append(groups[id-1], index)
					indexToGroup[index] = id
				}
				groups[oldId-1] = nil
			}
		}
	}
	counter := make([]map[int]int, len(groups))
	for i, group := range groups {
		if len(group) == 0 {
			continue
		}
		if counter[i] == nil {
			counter[i] = map[int]int{}
		}
		for _, index := range group {
			if c, ok := counter[i][source[index]]; ok {
				counter[i][source[index]] = c + 1
			} else {
				counter[i][source[index]] = 1
			}
		}
	}
	var out int
	for i := 0; i < n; i++ {
		id := indexToGroup[i]
		if id == 0 {
			if source[i] != target[i] {
				out++
			}
			continue
		}
		if c, ok := counter[id-1][target[i]]; ok {
			if c > 0 {
				counter[id-1][target[i]] = c - 1
			} else {
				out++
			}
		} else {
			out++
		}
	}
	return out
}
