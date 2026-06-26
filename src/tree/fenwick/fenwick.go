package fenwick

// 适合解前面有多少个符合条件的问题
type FenwickTree struct {
	nodes []int
}

func NewFenwickTree(n int) FenwickTree {
	return FenwickTree{
		nodes: make([]int, n+1),
	}
}

// 最右边的1，每个nodes[i]，表示[i - lowbit(i) + 1, i]区间的和
func lsb(n int) int {
	return n & (-n)
}

// j = i+lowbit(i)，表示寻找范围更大的且包含nodes[i]的nodes[j]
func (self *FenwickTree) Add(i int, v int) {
	i++
	for i < len(self.nodes) {
		self.nodes[i] += v
		i += lsb(i)
	}
}

// j = i-lowbit(i)，逐级将范围加起来，就是[0,i]区间的和
func (self *FenwickTree) Query(i int) int {
	i++
	var ret int
	for i > 0 {
		ret += self.nodes[i]
		i -= lsb(i)
	}
	return ret
}
