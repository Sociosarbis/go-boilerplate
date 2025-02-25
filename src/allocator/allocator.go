package allocator

type node struct {
	prev  *node
	next  *node
	start int
	end   int
	empty bool
}

type Allocator struct {
	root *node
	ids  map[int][]*node
}

func Constructor(n int) Allocator {
	return Allocator{
		root: &node{start: 0, end: n, empty: true},
		ids:  map[int][]*node{},
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	cur := this.root
	for cur != nil {
		if cur.empty && cur.end-cur.start >= size {
			cur.empty = false
			if cur.end-cur.start > size {
				newNode := &node{
					prev:  cur,
					next:  cur.next,
					start: cur.start + size,
					end:   cur.end,
					empty: true,
				}
				cur.end = cur.start + size
				if cur.next != nil {
					cur.next.prev = newNode
				}
				cur.next = newNode
			}

			this.ids[mID] = append(this.ids[mID], cur)
			return cur.start
		}
		cur = cur.next
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	var ret int
	for _, node := range this.ids[mID] {
		cur := node
		ret += node.end - node.start
		node.empty = true
		if node.prev != nil && node.prev.empty {
			cur = node.prev
			cur.end = node.end
			cur.next = node.next
			if cur.next != nil {
				cur.next.prev = cur
			}
		}
		if cur.next != nil && cur.next.empty {
			cur.end = cur.next.end
			cur.next = cur.next.next
			if cur.next != nil {
				cur.next.prev = cur
			}
		}
	}
	delete(this.ids, mID)
	return ret
}
