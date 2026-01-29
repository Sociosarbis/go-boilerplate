package cost

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type bfsItem struct {
	c int64
	i int
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	m := [26][26]int64{}
	g := [26][][2]int{}
	for i := 0; i < len(original); i++ {
		m[original[i]-'a'][changed[i]-'a'] = int64(cost[i])
		g[original[i]-'a'] = append(g[original[i]-'a'], [2]int{int(changed[i] - 'a'), cost[i]})
	}
	for i := 0; i < 26; i++ {
		hp := h.New([]bfsItem{{0, i}}, func(a, b bfsItem) bool {
			return a.c < b.c
		})
		for hp.Len() > 0 {
			top := heap.Pop(&hp).(bfsItem)
			n1, c1 := top.i, top.c
			if c1 > m[i][n1] {
				continue
			}
			for _, item := range g[n1] {
				n2, c2 := item[0], item[1]
				if i != n2 && (m[i][n2] == 0 || c1+int64(c2) <= m[i][n2]) {
					m[i][n2] = c1 + int64(c2)
					heap.Push(&hp, bfsItem{
						c: c1 + int64(c2),
						i: n2,
					})
				}
			}
		}
	}
	var out int64
	n := len(source)
	for i := 0; i < n; i++ {
		if source[i] == target[i] {
			continue
		}
		if m[source[i]-'a'][target[i]-'a'] == 0 {
			return -1
		}
		out += m[source[i]-'a'][target[i]-'a']
	}
	return out
}
