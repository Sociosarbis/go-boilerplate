package window

type node struct {
	val int
	i   int
}

type heap struct {
	arr []*node
}

func (n *node) GetParent(h *heap) *node {
	if n.i == 0 {
		return nil
	}
	return h.arr[(n.i-1)>>1]
}

func (n *node) GetLeft(h *heap) *node {
	i := (n.i << 1) + 1
	if i < len(h.arr) {
		return h.arr[i]
	}
	return nil
}

func (n *node) GetRight(h *heap) *node {
	i := (n.i << 1) + 2
	if i < len(h.arr) {
		return h.arr[i]
	}
	return nil
}

func (h *heap) Push(val int) *node {
	newNode := node{
		val,
		len(h.arr),
	}
	h.arr = append(h.arr, &newNode)
	h.BubbleUp(&newNode)
	return &newNode
}

func (h *heap) Pop(n *node) *node {
	if len(h.arr) != 0 {
		t := h.arr[len(h.arr)-1]
		h.Swap(n, t)
		h.arr = h.arr[:len(h.arr)-1]
		h.BubbleUp(t)
		h.SinkDown(t)
		return n
	}
	return nil
}

func (h *heap) BubbleUp(n *node) {
	p := n.GetParent(h)
	if p != nil {
		if p.val < n.val {
			h.Swap(p, n)
			h.BubbleUp(n)
		}
	}
}

func (h *heap) SinkDown(n *node) {
	l := n.GetLeft(h)
	r := n.GetRight(h)
	if l != nil && (r == nil || l.val >= r.val) && n.val < l.val {
		h.Swap(n, l)
		h.SinkDown(n)
	} else if r != nil && n.val < r.val {
		h.Swap(n, r)
		h.SinkDown(n)
	}
}

func (h *heap) Swap(a *node, b *node) {
	a.i, b.i = b.i, a.i
	h.arr[b.i] = b
	h.arr[a.i] = a
}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	h := heap{}
	win := []*node{}
	ret := []int{}
	for i := 0; i < k; i++ {
		win = append(win, h.Push(nums[i]))
	}
	ret = append(ret, h.arr[0].val)
	for i := k; i < len(nums); i++ {
		h.Pop(win[0])
		win = append(win[1:], h.Push(nums[i]))
		ret = append(ret, h.arr[0].val)
	}
	return ret
}
