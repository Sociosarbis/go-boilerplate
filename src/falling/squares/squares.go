package squares

type node struct {
	l      int
	r      int
	height int
	left   *node
	right  *node
}

func mergeNode(n1 *node, n2 *node) *node {
	height := n1.height
	if n2.height > height {
		height = n2.height
	}

	return &node{
		n1.l,
		n2.r,
		height,
		n1,
		n2,
	}
}

func traverse(n *node, l int, r int, height int) *node {
	if n.r <= l {
		return mergeNode(n, &node{
			l,
			r,
			height,
			nil,
			nil,
		})
	} else if n.l >= r {
		return mergeNode(&node{
			l,
			r,
			height,
			nil,
			nil,
		}, n)
	}

	// 超过左边
	if n.l >= l {
		// 超过右边
		if n.r <= r {
			return &node{
				l,
				r,
				height,
				nil,
				nil,
			}
		} else {
			if n.left == nil {
				nextLeft := &node{
					l,
					r,
					height,
					nil,
					nil,
				}

				nextRight := &node{
					r,
					n.r,
					n.height,
					nil,
					nil,
				}

				return mergeNode(nextLeft, nextRight)
			} else {
				// 如果没有跟右节点相交
				if n.right.l >= r {
					n.left = traverse(n.left, l, r, height)
				} else {
					n.left = traverse(n.left, l, n.right.l, height)
					n.right = traverse(n.right, n.right.l, r, height)
				}
			}
		}
	} else {
		if n.left == nil {
			// 如果不达右界，则分成3个
			if n.r > r {
				nextLeft := &node{
					n.l,
					l,
					n.height,
					nil,
					nil,
				}

				nextMid := &node{
					l,
					r,
					height,
					nil,
					nil,
				}

				nextRight := &node{
					r,
					n.r,
					n.height,
					nil,
					nil,
				}

				return mergeNode(mergeNode(nextLeft, nextMid), nextRight)
			} else {
				nextLeft := &node{
					n.l,
					l,
					n.height,
					nil,
					nil,
				}

				nextRight := &node{
					l,
					r,
					height,
					nil,
					nil,
				}
				return mergeNode(nextLeft, nextRight)
			}
		} else {
			// 如果与左节点相交
			if n.left.r > l {
				// 如果没有与右节点相交
				if n.right.l >= r {
					n.left = traverse(n.left, l, r, height)
				} else {
					n.left = traverse(n.left, l, n.right.l, height)
					n.right = traverse(n.right, n.right.l, r, height)
				}
			} else {
				// 如果没有与右节点相交
				if n.right.l >= r {
					n.left = mergeNode(n.left, &node{
						l,
						r,
						height,
						nil,
						nil,
					})
				} else {
					n.right = traverse(n.right, l, r, height)
				}
			}
		}
	}
	if n.l > n.left.l {
		n.l = n.left.l
	}

	if n.r < n.right.r {
		n.r = n.right.r
	}
	if n.left.height > n.height {
		n.height = n.left.height
	}
	if n.right.height > n.height {
		n.height = n.right.height
	}
	return n
}

func query(n *node, l int, r int) int {
	if n.r <= l {
		return 0
	} else if n.l >= r {
		return 0
	}

	if n.left == nil {
		return n.height
	}

	temp1 := query(n.left, l, r)
	temp2 := query(n.right, l, r)

	if temp1 > temp2 {
		return temp1
	} else {
		return temp2
	}
}

func fallingSquares(positions [][]int) []int {
	ret := make([]int, len(positions))
	ret[0] = positions[0][1]
	root := &node{
		positions[0][0],
		positions[0][0] + positions[0][1],
		positions[0][1],
		nil,
		nil,
	}

	for i := 1; i < len(positions); i++ {
		l := positions[i][0]
		r := l + positions[i][1]
		maxHeight := query(root, l, r)
		root = traverse(root, l, r, maxHeight+positions[i][1])
		ret[i] = root.height
	}
	return ret
}
