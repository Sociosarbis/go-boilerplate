package balanced2

type node struct {
	left  *node
	right *node
	l     int
	r     int
	min   int
	max   int
	diff  int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (self *node) merge(other *node) *node {
	return &node{
		l:     self.l,
		r:     other.r,
		min:   min(self.min, other.min),
		max:   max(self.max, other.max),
		left:  self,
		right: other,
	}
}

func (self *node) isLeaf() bool {
	return self.l == self.r
}

func (self *node) search(l, r int) int {
	if self.l > r || self.r < l || self.min > 0 || self.max < 0 {
		return -1
	}
	if !self.isLeaf() {
		self.down()
		res := self.right.search(l, r)
		if res != -1 {
			return res
		}
		res = self.left.search(l, r)
		if res != -1 {
			return res
		}
	}
	return self.l
}

func (self *node) refresh() {
	if self.isLeaf() {
		return
	}
	self.min = min(self.left.min, self.right.min)
	self.max = max(self.left.max, self.right.max)
}

func (self *node) down() {
	if self.diff == 0 || self.isLeaf() {
		return
	}
	self.left.update(self.l, self.r, self.diff)
	self.right.update(self.l, self.r, self.diff)
	self.diff = 0
}

func (self *node) update(l, r, diff int) {
	if self.l > r || self.r < l {
		return
	}
	if self.l >= l && self.r <= r {
		if self.l != self.r {
			self.diff += diff
		}
		self.min += diff
		self.max += diff
	} else {
		self.down()
		self.left.update(l, r, diff)
		self.right.update(l, r, diff)
		self.refresh()
	}
}

func longestBalanced(nums []int) int {
	n := len(nums)
	nodes := make([]*node, n)
	numToIndices := make(map[int][]int, n)
	visited := make(map[int]bool, n)
	var temp int
	for i, num := range nums {
		if _, ok := visited[num]; !ok {
			visited[num] = true
			if num&1 == 1 {
				temp++
			} else {
				temp--
			}
		}
		nodes[i] = &node{
			l:   i,
			r:   i,
			min: temp,
			max: temp,
		}
	}
	c := n
	for c != 1 {
		var index int
		for i := 0; i < c; i += 2 {
			if i+1 < c {
				nodes[index] = nodes[i].merge(nodes[i+1])
			} else {
				nodes[index] = nodes[i]
			}
			index++
		}
		c = index
	}
	for i := n - 1; i >= 0; i-- {
		if indices, ok := numToIndices[nums[i]]; ok {
			numToIndices[nums[i]] = append(indices, i)
		} else {
			numToIndices[nums[i]] = []int{i}
		}
	}
	root := nodes[0]
	out := root.search(0, n-1) + 1
	end := n - 2
	for i := 0; i < end; i++ {
		var diff int
		if nums[i]&1 == 1 {
			diff = -1
		} else {
			diff = 1
		}
		next := n
		if indices, ok := numToIndices[nums[i]]; ok {
			c := len(indices)
			if c != 0 {
				numToIndices[nums[i]] = indices[:c-1]
				if c > 1 {
					next = indices[c-2]
				}
			}
		}
		root.update(i+1, next-1, diff)
		r := root.search(i+1, n-1)
		if r-i > out {
			out = r - i
		}
	}
	return out
}
