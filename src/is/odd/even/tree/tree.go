package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	isOdd := 1
	bfs := make([]*TreeNode, 0)
	bfs = append(bfs, root)
	i := 0
	for i < len(bfs) {
		next_i := len(bfs)
		prev_val := 0
		if isOdd == 0 {
			prev_val = 1000000
		}
		for ; i < next_i; i += 1 {
			if (bfs[i].Val & 1) != isOdd {
				return false
			}
			if isOdd == 1 {
				if bfs[i].Val <= prev_val {
					return false
				}
			} else {
				if bfs[i].Val >= prev_val {
					return false
				}
			}
			prev_val = bfs[i].Val
			if bfs[i].Left != nil {
				bfs = append(bfs, bfs[i].Left)
			}
			if bfs[i].Right != nil {
				bfs = append(bfs, bfs[i].Right)
			}
		}
		i = next_i
		isOdd = 1 - isOdd
	}
	return true
}
