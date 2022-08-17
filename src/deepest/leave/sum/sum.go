package sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	bfs := []*TreeNode{root}
	for len(bfs) != 0 {
		n := len(bfs)
		for i := 0; i < n; i++ {
			if bfs[i].Left != nil {
				bfs = append(bfs, bfs[i].Left)
			}
			if bfs[i].Right != nil {
				bfs = append(bfs, bfs[i].Right)
			}
		}
		if n == len(bfs) {
			ret := 0
			for i := 0; i < n; i++ {
				ret += bfs[i].Val
			}
			return ret
		}
		bfs = bfs[n:]
	}
	return 0
}
