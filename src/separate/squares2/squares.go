package squares2

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type node struct {
	start int
	end   int
	value int
	// 快捷标签，表示已被覆盖
	count int
	left  *node
	right *node
}

func build(xl []int, l, r int) *node {
	out := &node{
		start: xl[l],
		end:   xl[r],
	}
	if l+1 == r {
		return out
	}
	mid := (l + r) / 2
	out.left = build(xl, l, mid)
	out.right = build(xl, mid, r)
	return out
}

func (self *node) isLeaf() bool {
	return self.left == nil && self.right == nil
}

func (self *node) getValue() int {
	if self.count > 0 {
		return self.end - self.start
	} else {
		if self.isLeaf() {
			return 0
		} else {
			return self.left.value + self.right.value
		}
	}
}

func (self *node) updateValue() {
	self.value = self.getValue()
}

func (self *node) change(start, end, delta int) {
	if self.end <= start || self.start >= end {
		return
	}
	if self.start >= start && self.end <= end {
		self.count += delta
		self.updateValue()
		return
	}
	self.left.change(start, end, delta)
	self.right.change(start, end, delta)
	self.updateValue()
}

type prefixItem struct {
	value int64
	end   int
}

type empty struct{}

func separateSquares(squares [][]int) float64 {
	n := len(squares)
	in, out := make([]int, n), make([]int, n)
	// 收集所有的x节点
	xs := make(map[int]empty, n)
	for _, square := range squares {
		xs[square[0]], xs[square[0]+square[2]] = empty{}, empty{}
	}
	xl := make([]int, 0, len(xs))
	for x := range xs {
		xl = append(xl, x)
	}
	sort.Ints(xl)
	// 提前build线段树结构
	root := build(xl, 0, len(xl)-1)
	for i := 0; i < n; i++ {
		in[i], out[i] = i, i
	}
	sort.Slice(in, func(i, j int) bool {
		return squares[in[i]][1] < squares[in[j]][1]
	})
	sort.Slice(out, func(i, j int) bool {
		return squares[out[i]][1]+squares[out[i]][2] < squares[out[j]][1]+squares[out[j]][2]
	})
	var i, j int
	root.change(squares[in[i]][0], squares[in[i]][0]+squares[in[i]][2], 1)
	i++
	for ; i < len(in); i++ {
		if squares[in[i]][1] == squares[in[i-1]][1] {
			root.change(squares[in[i]][0], squares[in[i]][0]+squares[in[i]][2], 1)
		} else {
			break
		}
	}
	prevStart := squares[in[0]][1]
	prefixSum := []prefixItem{
		{0, prevStart},
	}
	for i < len(in) || j < len(out) {
		var start int
		if j < len(out) {
			if i < len(in) {
				start = min(squares[in[i]][1], squares[out[j]][1]+squares[out[j]][2])
			} else {
				start = squares[out[j]][1] + squares[out[j]][2]
			}
		} else {
			start = squares[in[i]][1]
		}
		prefixSum = append(prefixSum, prefixItem{prefixSum[len(prefixSum)-1].value + int64(start-prevStart)*int64(root.getValue()), start})
		for ; j < len(out); j++ {
			if squares[out[j]][1]+squares[out[j]][2] == start {
				root.change(squares[out[j]][0], squares[out[j]][0]+squares[out[j]][2], -1)
			} else {
				break
			}
		}
		for ; i < len(in); i++ {
			if squares[in[i]][1] == start {
				root.change(squares[in[i]][0], squares[in[i]][0]+squares[in[i]][2], 1)
			} else {
				break
			}
		}
		prevStart = start
	}
	var l int
	r := len(prefixSum) - 1
	half := float64(prefixSum[r].value) / 2
	for l <= r {
		mid := (l + r) / 2
		d := float64(prefixSum[mid].value)
		if d < half {
			l = mid + 1
		} else {
			if mid > 0 && float64(prefixSum[mid-1].value) >= half {
				r = mid - 1
			} else {
				current := prefixSum[mid]
				prev := prefixSum[mid-1]
				return float64(prev.end) + float64(current.end-prev.end)*(half-float64(prev.value))/(float64(current.value)-float64(prev.value))
			}
		}
	}
	return 0
}
