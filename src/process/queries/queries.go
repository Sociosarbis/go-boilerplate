package queries

import "sort"

func processQueries(c int, connections [][]int, queries [][]int) []int {
	idToGroup := make([]int, c+1)
	groups := make([][]int, c+1)
	for i := 1; i <= c; i++ {
		idToGroup[i] = i
		groups[i] = append(groups[i], i)
	}
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		if idToGroup[a] == idToGroup[b] {
			continue
		}
		g1 := groups[idToGroup[a]]
		g2 := groups[idToGroup[b]]
		if len(g1) >= len(g2) {
			groups[idToGroup[a]] = append(g1, g2...)
			groups[idToGroup[b]] = nil
			for _, it := range g2 {
				idToGroup[it] = idToGroup[a]
			}
		} else {
			groups[idToGroup[b]] = append(g2, g1...)
			groups[idToGroup[a]] = nil
			for _, it := range g1 {
				idToGroup[it] = idToGroup[b]
			}
		}

	}
	for _, group := range groups {
		if len(group) > 1 {
			sort.Slice(group, func(i, j int) bool {
				return group[i] > group[j]
			})
		}
	}
	off := make([]bool, c+1)
	out := make([]int, 0, len(queries))
	for _, query := range queries {
		if query[0] == 2 {
			off[query[1]] = true
		} else {
			if !off[query[1]] {
				out = append(out, query[1])
			} else {
				index := idToGroup[query[1]]
				n := len(groups[index])
				i := n - 1
				for ; i >= 0; i-- {
					if !off[groups[index][i]] {
						break
					}
				}
				groups[index] = groups[index][:i+1]
				if i == -1 {
					out = append(out, -1)
				} else {
					out = append(out, groups[index][i])
				}
			}
		}
	}
	return out
}
