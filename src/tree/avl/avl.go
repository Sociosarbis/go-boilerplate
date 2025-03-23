package avl

type Comparable[T any] interface {
	gt(other T) bool
	lt(other T) bool
	eq(other T) bool
}

type Integer int

func (self Integer) lt(other Integer) bool {
	return self < other
}

func (self Integer) gt(other Integer) bool {
	return self > other
}

func (self Integer) eq(other Integer) bool {
	return self == other
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type AvlNode[T Comparable[T]] struct {
	value  T
	left   *AvlNode[T]
	right  *AvlNode[T]
	height int
}

func NewAvlNode[T Comparable[T]](value T) *AvlNode[T] {
	return &AvlNode[T]{
		value:  value,
		height: 1,
	}
}

func (self *AvlNode[T]) GetValue() T {
	return self.value
}

func (self *AvlNode[T]) getHeight() int {
	if self == nil {
		return 0
	}
	return self.height
}

func (self *AvlNode[T]) updateHeight() {
	self.height = 1 + max(self.left.getHeight(), self.right.getHeight())
}

func (self *AvlNode[T]) getBalance() int {
	if self == nil {
		return 0
	}
	return self.left.getHeight() - self.right.getHeight()
}

func (self *AvlNode[T]) leftRotate() *AvlNode[T] {
	r := self.right
	rl := r.left
	r.left = self
	self.right = rl
	self.updateHeight()
	r.updateHeight()
	return r
}

func (self *AvlNode[T]) rightRotate() *AvlNode[T] {
	l := self.left
	lr := l.right
	l.right = self
	self.left = lr
	self.updateHeight()
	l.updateHeight()
	return l
}

func (self *AvlNode[T]) getMinValueNode() *AvlNode[T] {
	if self == nil || self.left == nil {
		return self
	}
	return self.left.getMinValueNode()
}

func (self *AvlNode[T]) balance() *AvlNode[T] {
	self.updateHeight()
	balanceFactor := self.getBalance()
	if balanceFactor > 1 {
		if self.left.getBalance() >= 0 {
			return self.rightRotate()
		} else {
			self.left = self.left.leftRotate()
			return self.rightRotate()
		}
	}
	if balanceFactor < -1 {
		if self.right.getBalance() <= 0 {
			return self.leftRotate()
		} else {
			self.right = self.right.rightRotate()
			return self.leftRotate()
		}
	}
	return self
}

func (self *AvlNode[T]) Insert(value T) *AvlNode[T] {
	if self == nil {
		return NewAvlNode[T](value)
	}
	if value.lt(self.value) {
		self.left = self.left.Insert(value)
	} else {
		self.right = self.right.Insert(value)
	}
	return self.balance()
}

func (self *AvlNode[T]) Delete(value T) *AvlNode[T] {
	if self == nil {
		return self
	}
	if value.lt(self.value) {
		self.left = self.left.Delete(value)
	} else if value.gt(self.value) {
		self.right = self.right.Delete(value)
	} else {
		if self.left == nil {
			return self.right
		} else if self.right == nil {
			return self.left
		}
		rm := self.right.getMinValueNode()
		self.value = rm.value
		self.right = self.right.Delete(rm.value)
	}
	return self.balance()
}

func (self *AvlNode[T]) LowerBound(value T) *avlNodeIterator[T] {
	if self == nil {
		return nil
	}
	cur := self
	path := []*AvlNode[T]{}
	for cur != nil {
		if cur.value.lt(value) {
			path = append(path, cur)
			cur = cur.right
		} else {
			path = append(path, cur)
			cur = cur.left
		}
	}
	n := len(path)
	if n != 0 {
		ret := &avlNodeIterator[T]{
			path: path[:n-1],
			cur:  path[n-1],
		}
		for ; ret.cur != nil && ret.cur.GetValue().lt(value); ret.Next() {
		}
		return ret
	}
	return nil
}

type avlNodeIterator[T Comparable[T]] struct {
	path []*AvlNode[T]
	cur  *AvlNode[T]
}

func (self *avlNodeIterator[T]) GetCur() *AvlNode[T] {
	if self == nil {
		return nil
	}
	return self.cur
}

func (self *avlNodeIterator[T]) Next() *AvlNode[T] {
	if self == nil || self.cur == nil {
		return nil
	}
	if self.cur.right != nil {
		self.path = append(self.path, self.cur)
		self.cur = self.cur.right
		for self.cur.left != nil {
			self.path = append(self.path, self.cur)
			self.cur = self.cur.left
		}
		return self.cur
	}
	n := len(self.path)
	for n != 0 {
		p := self.path[n-1]
		self.path = self.path[:n-1]
		if self.cur == p.left {
			self.cur = p
			return p
		}
		self.cur = p
		n = len(self.path)
	}
	self.cur = nil
	return nil
}
