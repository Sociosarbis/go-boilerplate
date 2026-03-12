package stability

import "sort"

type dsu struct {
	parent []int
}

func (self *dsu) find(x int) int {
	if self.parent[x] == x {
		return x
	}
	self.parent[x] = self.find(self.parent[x])
	return self.parent[x]
}

func (self *dsu) join(x, y int) {
	px, py := self.find(x), self.find(y)
	if px != py {
		self.parent[px] = py
	}
}

func maxStability(n int, edges [][]int, k int) int {
	if len(edges) < n-1 {
		return -1
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] > edges[j][2]
	})
	currentEdges := make([][]int, 0, len(edges))
	for _, edge := range edges {
		if edge[3] == 1 {
			currentEdges = append(currentEdges, edge)
		}
	}
	if len(currentEdges) > n-1 {
		return -1
	}
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	unionInit := dsu{
		parent,
	}
	var out int = 2 * 1e5
	selectedInit := len(currentEdges)
	if selectedInit > 0 {
		for _, edge := range currentEdges {
			if unionInit.find(edge[0]) == unionInit.find(edge[1]) {
				return -1
			}
			unionInit.join(edge[0], edge[1])
		}
		out = currentEdges[len(currentEdges)-1][2]
	}
	var i int
	for _, edge := range edges {
		if edge[3] == 0 {
			if i < len(currentEdges) {
				currentEdges[i] = edge
			} else {
				currentEdges = append(currentEdges, edge)
			}
			i++
		}
	}
	currentEdges = currentEdges[:i]
	var l int
	r := out
	out = -1
	for l <= r {
		mid := l + ((r - l) >> 1)
		parent := make([]int, n)
		copy(parent, unionInit.parent)
		union := dsu{
			parent,
		}
		count := selectedInit
		remain := k
		// 利用不成环 = 都在同一个集合里 且 边为n - 1的性质
		for _, edge := range currentEdges {
			u, v, s := edge[0], edge[1], edge[2]
			if union.find(u) == union.find(v) {
				continue
			}
			union.join(u, v)
			if s >= mid {
				count++
			} else if remain > 0 && s*2 >= mid {
				remain--
				count++
			} else {
				break
			}
			if count >= n-1 {
				break
			}
		}
		if count >= n-1 {
			out = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return out
}
