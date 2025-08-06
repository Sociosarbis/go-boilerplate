package fruits2

type node struct {
	left  *node
	right *node
	start int
	end   int
	v     int
}

func (this *node) merge(other *node) *node {
	var v int
	if this.v > other.v {
		v = this.v
	} else {
		v = other.v
	}
	return &node{
		left:  this,
		right: other,
		start: this.start,
		end:   this.end,
		v:     v,
	}
}

func (this *node) update() {
	var v int
	if this.left != nil {
		v = this.left.v
	}
	if this.right != nil {
		if this.right.v > v {
			v = this.right.v
		}
	}

	this.v = v
}

func (this *node) remove(value int) bool {
	if this.v >= value {
		if this.left != nil && this.left.v >= value {
			this.left.remove(value)
			this.update()
		} else if this.right != nil && this.right.v >= value {
			this.right.remove(value)
			this.update()
		} else {
			this.v = 0
		}
		return true
	}
	return false
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	nodes := make([]*node, 0, len(baskets))
	for i, basket := range baskets {
		nodes = append(nodes, &node{
			start: i,
			end:   i,
			v:     basket,
		})
	}
	for len(nodes) > 1 {
		var j int
		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				nodes[j] = nodes[i].merge(nodes[i+1])
			} else {
				nodes[j] = nodes[i]
			}
			j++
		}
		nodes = nodes[:j]
	}
	root := nodes[0]
	var out int
	for _, fruit := range fruits {
		if !root.remove(fruit) {
			out++
		}
	}
	return out
}
