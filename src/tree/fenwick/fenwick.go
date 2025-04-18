package fenwick

type FenwickTree struct {
	nodes []int
}

func NewFenwickTree(n int) FenwickTree {
	return FenwickTree{
		nodes: make([]int, n+1),
	}
}

func lsb(n int) int {
	return n & (-n)
}

func (self *FenwickTree) Add(i int, v int) {
	i++
	for i < len(self.nodes) {
		self.nodes[i] += v
		i += lsb(i)
	}
}

func (self *FenwickTree) Query(i int) int {
	i++
	var ret int
	for i > 0 {
		ret += self.nodes[i]
		i -= lsb(i)
	}
	return ret
}
