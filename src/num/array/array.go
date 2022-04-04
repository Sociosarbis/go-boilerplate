package array

type node struct {
	val   int
	i     int
	j     int
	left  *node
	right *node
}

type NumArray struct {
	root *node
}

func Constructor(nums []int) NumArray {
	nodes := make([]*node, len(nums))

	for i, num := range nums {
		nodes[i] = &node{
			num,
			i,
			i,
			nil,
			nil,
		}
	}

	for len(nodes) > 1 {
		j := 0
		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				nodes[j] = &node{
					nodes[i].val + nodes[i+1].val,
					nodes[i].i,
					nodes[i+1].j,
					nodes[i],
					nodes[i+1],
				}
			} else {
				nodes[j] = nodes[i]
			}
			j++
		}
		nodes = nodes[:j]
	}
	return NumArray{
		nodes[0],
	}
}

func (this *NumArray) Update(index int, val int) {
	updateDfs(this.root, index, val)
}

func updateDfs(node *node, index int, val int) {
	if node.i == node.j {
		node.val = val
	} else {
		if node.left.i <= index && node.left.j >= index {
			updateDfs(node.left, index, val)
		} else {
			updateDfs(node.right, index, val)
		}
		node.val = node.left.val + node.right.val
	}

}

func (this *NumArray) SumRange(left int, right int) int {
	return sumRangeDfs(this.root, left, right)
}

func sumRangeDfs(node *node, left int, right int) int {
	if node.i == left && node.j == right {
		return node.val
	}
	leftSum := 0
	rightSum := 0
	if node.left.j >= left {
		l := left
		r := right
		if node.left.j < r {
			r = node.left.j
		}
		leftSum = sumRangeDfs(node.left, l, r)
	}

	if node.right.i <= right {
		l := left
		r := right
		if node.right.i > left {
			l = node.right.i
		}
		rightSum = sumRangeDfs(node.right, l, r)
	}
	return leftSum + rightSum
}
