package order2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	bfs := []*TreeNode{root}
	for len(bfs) != 0 {
		n := len(bfs)
		temp := []int{}
		for i := 0; i < n; i++ {
			temp = append(temp, bfs[i].Val)
			if bfs[i].Left != nil {
				bfs = append(bfs, bfs[i].Left)
			}

			if bfs[i].Right != nil {
				bfs = append(bfs, bfs[i].Right)
			}
		}
		ret = append(ret, temp)
		bfs = bfs[n:]
	}
	return ret
}
