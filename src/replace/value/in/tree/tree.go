package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	bfs := []*TreeNode{root}
	for len(bfs) != 0 {
		sum := 0
		n := len(bfs)
		for i := 0; i < n; i++ {
			current := bfs[i]
			if current.Left != nil {
				sum += current.Left.Val
			}
			if current.Right != nil {
				sum += current.Right.Val
			}
		}
		for i := 0; i < n; i++ {
			current := bfs[i]
			total := 0
			if current.Left != nil {
				total += current.Left.Val
			}
			if current.Right != nil {
				total += current.Right.Val
			}
			if current.Left != nil {
				current.Left.Val = sum - total
				bfs = append(bfs, current.Left)
			}
			if current.Right != nil {
				current.Right.Val = sum - total
				bfs = append(bfs, current.Right)
			}
		}
		bfs = bfs[n:]
	}
	return root
}
