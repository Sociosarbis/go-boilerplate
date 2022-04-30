package construct

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func dfs(grid *[][]int, i, j, n int) *Node {
	node := &Node{}

	if (*grid)[i][j] == 1 {
		node.Val = true
	}

	if n != 1 {
		halfN := n >> 1
		for k := i; k < i+n; k += halfN {
			for l := j; l < j+n; l += halfN {
				child := dfs(grid, k, l, halfN)

				if k == i {
					if l == j {
						node.TopLeft = child
					} else {
						node.TopRight = child
					}
				} else {
					if l == j {
						node.BottomLeft = child
					} else {
						node.BottomRight = child
					}
				}
			}
		}

		if node.BottomLeft.IsLeaf && node.BottomRight.IsLeaf && node.TopLeft.IsLeaf && node.TopRight.IsLeaf && node.TopLeft.Val == node.TopRight.Val && node.TopLeft.Val == node.BottomRight.Val && node.TopLeft.Val == node.BottomLeft.Val {
			node.IsLeaf = true
			node.TopLeft = nil
			node.TopRight = nil
			node.BottomLeft = nil
			node.BottomRight = nil
		}
	} else {
		node.IsLeaf = true
	}

	return node
}

func construct(grid [][]int) *Node {
	return dfs(&grid, 0, 0, len(grid))
}
