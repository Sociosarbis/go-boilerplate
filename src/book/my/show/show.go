package show

type rng struct {
	start int
	end   int
}

func (self *rng) width() int {
	return self.end - self.start
}

func (self *rng) max(other *rng) rng {
	if self.width() > other.width() {
		return *self
	}
	return *other
}

type base struct {
	capacity    int64
	maxColRange rng
}

type rowNode struct {
	base
	rowRange rng
	left     *rowNode
	right    *rowNode
	row      *colNode
}

func min[T int | int64](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (self *rowNode) merge(other *rowNode) rowNode {
	left, right := self, other
	if self.rowRange.start > other.rowRange.start {
		left, right = right, left
	}
	return rowNode{
		base: base{
			capacity:    left.capacity + right.capacity,
			maxColRange: left.maxColRange.max(&right.maxColRange),
		},
		rowRange: rng{
			start: left.rowRange.start,
			end:   right.rowRange.end,
		},
		left:  left,
		right: right,
	}
}

func (self *rowNode) tryGather(k int, maxRow int) []int {
	if self.maxColRange.width() < k || self.rowRange.start > maxRow {
		return nil
	}
	if self.left != nil {
		r := self.left.tryGather(k, maxRow)
		if r != nil {
			if self.left.capacity == 0 {
				*self = *self.right
			} else {
				*self = self.left.merge(self.right)
			}
			return r
		}
	}
	if self.right != nil {
		r := self.right.tryGather(k, maxRow)
		if r != nil {
			if self.right.capacity == 0 {
				*self = *self.left
			} else {
				*self = self.left.merge(self.right)
			}
			return r
		}
	} else {
		c := self.row.tryGather(k)
		if c != -1 {
			self.capacity = self.row.capacity
			self.maxColRange = self.row.maxColRange
			return []int{self.rowRange.start, c}
		}
	}
	return nil
}

func (self *rowNode) getCapacity(maxRow int) int64 {
	if self.rowRange.start > maxRow {
		return 0
	}
	if self.rowRange.end <= maxRow+1 {
		return self.capacity
	}
	if self.left != nil {
		return self.left.getCapacity(maxRow) + self.right.getCapacity(maxRow)
	} else {
		return self.capacity
	}
}

func (self *rowNode) scatter(k int, maxRow int) {
	if self.rowRange.start > maxRow {
		return
	}
	if self.left != nil {
		oldCapacity := self.left.capacity
		self.left.scatter(k, maxRow)
		k -= int(oldCapacity - self.left.capacity)
	}

	if k != 0 {
		if self.right != nil {
			self.right.scatter(k, maxRow)
		} else {
			self.row.scatter(k)
		}
	}

	if self.left != nil {
		if self.left.capacity == 0 {
			*self = *self.right
		} else {
			if self.right.capacity == 0 {
				*self = *self.left
			} else {
				*self = self.left.merge(self.right)
			}
		}
	} else {
		self.capacity = self.row.capacity
		self.maxColRange = self.row.maxColRange
	}
}

type colNode struct {
	base
	colRange rng
	left     *colNode
	right    *colNode
}

func (self *colNode) merge(other *colNode) colNode {
	left, right := self, other
	if self.left.colRange.start > right.colRange.start {
		left, right = right, left
	}
	return colNode{
		base: base{
			capacity:    left.capacity + right.capacity,
			maxColRange: left.maxColRange.max(&right.maxColRange),
		},
		colRange: rng{
			start: left.colRange.start,
			end:   right.colRange.end,
		},
		left:  left,
		right: right,
	}
}

func (self *colNode) tryGather(k int) int {
	if self.capacity < int64(k) || self.maxColRange.width() < k {
		return -1
	}
	if self.left != nil {
		r := self.left.tryGather(k)
		if r != -1 {
			if self.left.capacity == 0 {
				*self = *self.right
			} else {
				*self = self.left.merge(self.right)
			}
			return r
		}
	}

	if self.right != nil {
		r := self.right.tryGather(k)
		if r != -1 {
			if self.right.capacity == 0 {
				*self = *self.left
			} else {
				*self = self.left.merge(self.right)
			}
			return r
		}
	} else {
		ret := self.colRange.start
		self.colRange.start += k
		self.capacity -= int64(k)
		self.maxColRange = self.colRange
		return ret
	}
	return -1
}

func (self *colNode) scatter(k int) {
	if self.left != nil {
		oldCapacity := self.left.capacity
		self.left.scatter(k)
		k -= int(oldCapacity - self.left.capacity)
	}

	if k != 0 {
		if self.right != nil {
			self.right.scatter(k)
		} else {
			d := min(int64(k), self.capacity)
			self.capacity -= d
			self.colRange.start += int(d)
			self.maxColRange = self.colRange
		}
	}

	if self.left != nil {
		if self.left.capacity == 0 {
			*self = *self.right
		} else {
			if self.right.capacity == 0 {
				*self = *self.left
			} else {
				*self = self.left.merge(self.right)
			}
		}
	}
}

type BookMyShow struct {
	root rowNode
}

func buildTree(start int, n int, m int) rowNode {
	if n == 1 {
		b := base{
			capacity: int64(m),
			maxColRange: rng{
				0,
				m,
			},
		}
		return rowNode{
			base: b,
			rowRange: rng{
				start: start,
				end:   start + n,
			},
			row: &colNode{base: b, colRange: b.maxColRange},
		}
	}
	n1 := n / 2
	n2 := n - n1
	left := buildTree(start, n1, m)
	right := buildTree(start+n1, n2, m)
	return left.merge(&right)
}

func Constructor(n int, m int) BookMyShow {
	return BookMyShow{
		root: buildTree(0, n, m),
	}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	return this.root.tryGather(k, maxRow)
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if this.root.getCapacity(maxRow) >= int64(k) {
		this.root.scatter(k, maxRow)
		return true
	}
	return false
}
