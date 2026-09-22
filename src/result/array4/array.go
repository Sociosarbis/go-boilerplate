package array4

type node struct {
	count []int
	next  []int
}

func newNode(k int) node {
	return node{
		count: make([]int, k),
		next:  make([]int, k),
	}
}

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func split(i, l, r int) (int, int, int, int, int, int) {
	li, ri := children(i)
	mid := (l + r) / 2
	ll, lr, rl, rr := l, mid, mid+1, r
	return li, ll, lr, ri, rl, rr
}

func buildTree(nodes []node, nums []int, i, l, r, k int) {
	if l == r {
		s := newNode(k)
		index := nums[l] % k
		s.count[index] = 1
		s.next[index] = 1
		nodes[i] = s
	} else {
		li, ll, lr, ri, rl, rr := split(i, l, r)
		buildTree(nodes, nums, li, ll, lr, k)
		buildTree(nodes, nums, ri, rl, rr, k)
		nodes[i] = newNode(k)

		bottomUp(nodes[i].count, nodes[i].next, nodes[li].count, nodes[li].next, nodes[ri].count, nodes[ri].next, k)
	}
}

func bottomUp(count, next, lc, ln, rc, rn []int, k int) {
	copy(count, lc)
	for i := 0; i < k; i++ {
		next[i] = 0
	}
	for a := 0; a < k; a++ {
		if ln[a] != 0 {
			for b := 0; b < k; b++ {
				index := (a * b) % k
				if rc[b] != 0 {
					count[index] += ln[a] * rc[b]
				}
				if rn[b] != 0 {
					next[index] += ln[a] * rn[b]
				}
			}
		}
	}
}

func update(nodes []node, i, l, r, index, value, k int) {
	if index < l || index > r {
		return
	}
	if l == r {
		for j := 0; j < k; j++ {
			nodes[i].count[j] = 0
			nodes[i].next[j] = 0
		}
		idx := value % k
		nodes[i].count[idx] = 1
		nodes[i].next[idx] = 1
	} else {
		li, ll, lr, ri, rl, rr := split(i, l, r)
		update(nodes, li, ll, lr, index, value, k)
		update(nodes, ri, rl, rr, index, value, k)
		bottomUp(nodes[i].count, nodes[i].next, nodes[li].count, nodes[li].next, nodes[ri].count, nodes[ri].next, k)
	}
}

func query(nodes []node, i, l, r, start, k int) ([]int, []int) {
	if start > r {
		return nil, nil
	}
	if start <= l {
		return nodes[i].count, nodes[i].next
	}
	li, ll, lr, ri, rl, rr := split(i, l, r)
	lc, ln := query(nodes, li, ll, lr, start, k)
	rc, rn := query(nodes, ri, rl, rr, start, k)
	if lc == nil {
		return rc, rn
	}
	next := make([]int, k)
	if rc == nil {
		return lc, next
	}
	count := make([]int, k)
	bottomUp(count, next, lc, ln, rc, rn, k)
	return count, next
}

func resultArray(nums []int, k int, queries [][]int) []int {
	n := len(nums)
	count := 1
	for n != 0 {
		n >>= 1
		count++
	}
	nodes := make([]node, 1<<count)
	n = len(nums)
	buildTree(nodes, nums, 0, 0, n-1, k)
	out := make([]int, len(queries))
	for i, q := range queries {
		index, value, start, x := q[0], q[1], q[2], q[3]
		update(nodes, 0, 0, n-1, index, value, k)
		res, _ := query(nodes, 0, 0, n-1, start, k)
		out[i] = res[x]
	}
	return out
}
