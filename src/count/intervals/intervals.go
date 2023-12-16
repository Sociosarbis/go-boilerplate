package intervals

type node struct {
	count int
	start int
	end   int
	left  *node
	right *node
}

type CountIntervals struct {
	root *node
}

func Constructor() CountIntervals {
	return CountIntervals{}
}

func (this *CountIntervals) Add(left int, right int) {
	temp := &node{
		count: right - left + 1,
		start: left,
		end:   right,
	}
	if this.root == nil {
		this.root = temp
		return
	}
	this.root = add(this.root, temp)
}

func add(a *node, b *node) *node {
	if a.end < b.start || b.end < a.start {
		if a.start > b.start {
			a, b = b, a
		}
		return &node{
			count: a.count + b.count,
			start: a.start,
			end:   b.end,
			left:  a,
			right: b,
		}
	}
	if a.left == nil {
		if a.start > b.start {
			a.start = b.start
		}
		if a.end < b.end {
			a.end = b.end
		}
		a.count = a.end - a.start + 1
		return a
	}

	if b.start < a.right.start {
		temp := *b
		node := &temp
		if node.end >= a.right.start {
			node.end = a.right.start - 1
		}
		node.count = node.end - node.start + 1
		a.left = add(a.left, node)
	}

	if b.end >= a.right.start {
		temp := *b
		node := &temp
		if node.start < a.right.start {
			node.start = a.right.start
		}
		node.count = node.end - node.start + 1
		a.right = add(a.right, node)

	}
	a.count = a.left.count + a.right.count
	a.start = a.left.start
	a.end = a.right.end

	if a.left.left == nil && a.right.left == nil && a.left.end+1 == a.right.start {
		a.left = nil
		a.right = nil
	}
	return a
}

func (this *CountIntervals) Count() int {
	if this.root == nil {
		return 0
	}
	return this.root.count
}
