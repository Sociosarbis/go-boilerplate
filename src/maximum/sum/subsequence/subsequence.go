package subsequence

const mod int = 1e9 + 7

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type node struct {
	parent *node
	left   *node
	right  *node
	values [4]int
}

func (n *node) update() {
	if n.left != nil {
		n.values[0] = max(n.left.values[0]+n.right.values[1], n.left.values[2]+n.right.values[0]) % mod
		n.values[1] = max(n.left.values[1]+max(n.right.values[0], n.right.values[1]), n.left.values[3]+n.right.values[0]) % mod
		n.values[2] = max(max(n.left.values[0], n.left.values[2])+n.right.values[2], n.left.values[0]+n.right.values[3]) % mod
		n.values[3] = max(n.left.values[1]+max(n.right.values[2], n.right.values[3]), n.left.values[3]+n.right.values[2]) % mod
	}
	if n.parent != nil {
		n.parent.update()
	}
}

func (n *node) merge(other *node) *node {
	parent := &node{
		left:  n,
		right: other,
	}
	n.parent = parent
	other.parent = parent
	parent.update()
	return parent
}

func (n *node) setValue(v int) {
	n.values[3] = v
	n.update()
}

func (n *node) max() int {
	return max(max(n.values[0], n.values[1]), max(n.values[2], n.values[3]))
}

func buildTree(nums []int, leaves []*node, start, end int) *node {
	if start == end {
		leaves[start] = &node{
			values: [4]int{0, 0, 0, nums[start]},
		}
		return leaves[start]
	}
	mid := (start + end) / 2
	return buildTree(nums, leaves, start, mid).merge(buildTree(nums, leaves, mid+1, end))
}

func maximumSumSubsequence(nums []int, queries [][]int) int {
	n := len(nums)
	leaves := make([]*node, n)
	root := buildTree(nums, leaves, 0, n-1)
	var ret int
	for _, q := range queries {
		i, v := q[0], q[1]
		leaves[i].setValue(v)
		ret = (ret + root.max()) % mod
	}
	return ret
}
