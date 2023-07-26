package query

type node struct {
	value   int
	left    int
	right   int
	lazyTag bool
}

func (n *node) count() int {
	return n.right - n.left + 1
}

func (n *node) flip() {
	n.lazyTag = !n.lazyTag
	n.value = n.count() - n.value
}

func (n *node) leftRightBound() int {
	return (n.left + n.right) / 2
}

func leftIndex(index int) int {
	return index*2 + 1
}

func rightIndex(index int) int {
	return (index + 1) * 2
}

func buildTree(tree *[]node, nums1 *[]int, index int, left int, right int) {
	n := &(*tree)[index]
	n.left = left
	n.right = right
	if left != right {
		mid := n.leftRightBound()
		buildTree(tree, nums1, leftIndex(index), left, mid)
		buildTree(tree, nums1, rightIndex(index), mid+1, right)
		(*tree)[index].value = (*tree)[leftIndex(index)].value + (*tree)[rightIndex(index)].value
	} else {
		(*tree)[index].value = (*nums1)[left]
	}
}

func pushDown(tree *[]node, index int) {
	n := &(*tree)[index]
	if n.lazyTag {
		(*tree)[leftIndex(index)].flip()
		(*tree)[rightIndex(index)].flip()
		n.lazyTag = false
	}
}

func add(tree *[]node, index int, left int, right int) {
	n := &(*tree)[index]
	if n.left >= left && n.right <= right {
		n.flip()
		return
	}
	pushDown(tree, index)
	if n.leftRightBound() >= left {
		add(tree, leftIndex(index), left, right)
	}

	if n.leftRightBound()+1 <= right {
		add(tree, rightIndex(index), left, right)
	}
	n.value = (*tree)[leftIndex(index)].value + (*tree)[rightIndex(index)].value
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := 1
	for n < len(nums1) {
		n *= 2
	}
	tree := make([]node, n*2-1)
	var temp int64
	for _, num := range nums2 {
		temp += int64(num)
	}
	buildTree(&tree, &nums1, 0, 0, len(nums1)-1)
	ret := []int64{}
	for _, q := range queries {
		if q[0] == 1 {
			add(&tree, 0, q[1], q[2])
		} else if q[0] == 2 {
			temp += int64(tree[0].value) * int64(q[1])
		} else {
			ret = append(ret, temp)
		}
	}
	return ret
}
