package results

import "sort"

// obstacle的位置
type set []int

func (self set) insert(value int) set {
	index := sort.Search(len(self), func(i int) bool {
		return (self)[i] > value
	})
	self = append(self, 0)
	copy(self[index+1:], self[index:])
	self[index] = value
	return self
}

func (self set) prevNext(value int) (int, int) {
	index := sort.Search(len(self), func(i int) bool {
		return (self)[i] >= value
	})
	if self[index] == value {
		return self[index-1], self[index+1]
	}
	return self[index-1], self[index]
}

// 范围内最大的到左侧obstacle的距离
type seg []int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (self seg) update(idx, val, p, l, r int) seg {
	if l == r {
		self[p] = val
		return self
	}
	mid := (l + r) >> 1
	if idx <= mid {
		self = self.update(idx, val, p<<1, l, mid)
	} else {
		self = self.update(idx, val, p<<1|1, mid+1, r)
	}
	self[p] = max(self[p<<1], self[p<<1|1])
	return self
}

func (self seg) query(R, p, l, r int) int {
	if R >= r {
		return self[p]
	}
	mid := (l + r) >> 1
	out := self.query(R, p<<1, l, mid)
	if mid+1 <= R {
		out = max(out, self.query(R, p<<1|1, mid+1, r))
	}
	return out
}

func getResults(queries [][]int) []bool {
	out := make([]bool, 0, len(queries))
	mx := min(50000, 3*len(queries)) + 1
	st := set(make([]int, 0, mx+1))
	st = append(st, 0, mx)
	sg := seg(make([]int, mx<<2))
	sg.update(mx, mx, 1, 0, mx)
	for _, query := range queries {
		if query[0] == 1 {
			prev, next := st.prevNext(query[1])
			sg = sg.update(next, next-query[1], 1, 0, mx)
			sg = sg.update(query[1], query[1]-prev, 1, 0, mx)
			st = st.insert(query[1])
		} else {
			prev, _ := st.prevNext(query[1])
			out = append(out, query[1]-prev >= query[2] || sg.query(prev, 1, 0, mx) >= query[2])
		}
	}
	return out
}
