package people

import "sort"

func processMeetings(meetings [][]int, known []bool, start, end int) {
	n := end - start + 1
	idToGroup := make(map[int]int, n)
	groupToIds := make([][]int, 0, n)
	for i := start; i <= end; i++ {
		x, y := meetings[i][0], meetings[i][1]
		if g1, ok := idToGroup[x]; ok {
			if g2, ok := idToGroup[y]; ok {
				if g1 != g2 {
					for _, id := range groupToIds[g2] {
						idToGroup[id] = g1
					}
					groupToIds[g1] = append(groupToIds[g1], groupToIds[g2]...)
					groupToIds[g2] = nil
				}
			} else {
				idToGroup[y] = g2
				groupToIds[g1] = append(groupToIds[g1], y)
			}
		} else {
			if g2, ok := idToGroup[y]; ok {
				idToGroup[x] = g2
				groupToIds[g2] = append(groupToIds[g2], x)
			} else {
				group := len(groupToIds)
				idToGroup[x], idToGroup[y] = group, group
				groupToIds = append(groupToIds, []int{x, y})
			}
		}
	}
	for _, ids := range groupToIds {
		if len(ids) != 0 {
			var found bool
			for i, id := range ids {
				if found {
					known[id] = true
				} else if known[id] {
					for j := 0; j < i; j++ {
						known[ids[j]] = true
					}
					found = true
				}
			}
		}
	}
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	known := make([]bool, n)
	known[0], known[firstPerson] = true, true
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	var start int
	m := len(meetings)
	for i := 1; i < m; i++ {
		if meetings[i][2] != meetings[start][2] {
			processMeetings(meetings, known, start, i-1)
			start = i
		}
	}
	processMeetings(meetings, known, start, m-1)
	out := make([]int, 0, n)
	for i, yes := range known {
		if yes {
			out = append(out, i)
		}
	}
	return out
}
