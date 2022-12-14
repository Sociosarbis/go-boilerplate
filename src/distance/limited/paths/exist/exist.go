package exist

import (
	"sort"
)

type edges [][]int

func (this edges) Len() int {
	return len(this)
}

func (this edges) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this edges) Less(i, j int) bool {
	return this[i][2] < this[j][2]
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Sort(edges(edgeList))
	for i, q := range queries {
		queries[i] = append(q, i)
	}
	sort.Sort(edges(queries))
	ret := make([]bool, len(queries))
	groups := map[int][]int{}
	nextGroupId := 1
	m := make([]int, n)
	i := 0
	for _, e := range edgeList {
		l, r, c := e[0], e[1], e[2]
		for i < len(queries) && queries[i][2] <= c {
			if m[queries[i][0]] != 0 && m[queries[i][1]] != 0 && m[queries[i][0]] == m[queries[i][1]] {
				ret[queries[i][3]] = true
			}
			i++
		}
		if i == len(queries) {
			break
		}
		if m[l] != 0 {
			if m[r] != 0 {
				if m[l] != m[r] {
					g1 := groups[m[l]]
					g2 := groups[m[r]]
					if len(g1) >= len(g2) {
						groups[m[l]] = append(g1, g2...)
						delete(groups, m[r])
						for _, index := range g2 {
							m[index] = m[l]
						}
					} else {
						groups[m[r]] = append(g2, g1...)
						delete(groups, m[l])
						for _, index := range g1 {
							m[index] = m[r]
						}
					}
				}
			} else {
				m[r] = m[l]
				groups[m[l]] = append(groups[m[l]], r)
			}
		} else {
			if m[r] != 0 {
				m[l] = m[r]
				groups[m[r]] = append(groups[m[r]], l)
			} else {
				m[l] = nextGroupId
				m[r] = nextGroupId
				groups[nextGroupId] = []int{l, r}
				nextGroupId++
			}
		}
		if m[queries[i][0]] != 0 && m[queries[i][1]] != 0 && m[queries[i][0]] == m[queries[i][1]] {
			ret[queries[i][3]] = true
		}
	}
	for ; i < len(queries); i++ {
		ret[queries[i][3]] = true
	}
	return ret
}
