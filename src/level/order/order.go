package order

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ret := [][]int{}

	if root == nil {
		return ret
	}

	bfs := []*Node{root}

	for len(bfs) != 0 {
		temp := []int{}
		nextBfs := []*Node{}

		for i := 0; i < len(bfs); i++ {
			temp = append(temp, bfs[i].Val)
			nextBfs = append(nextBfs, bfs[i].Children...)
		}
		ret = append(ret, temp)
		bfs = nextBfs
	}

	return ret
}
