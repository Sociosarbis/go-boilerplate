package sum

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	hp := h.New([]int64{}, func(a, b int64) bool {
		return a < b
	})

	bfs := []*TreeNode{root}
	n := len(bfs)
	for n != 0 {
		var sum int64
		for i := 0; i < n; i++ {
			node := bfs[i]
			sum += int64(node.Val)
			if node.Left != nil {
				bfs = append(bfs, node.Left)
			}
			if node.Right != nil {
				bfs = append(bfs, node.Right)
			}
		}
		heap.Push(&hp, sum)
		if hp.Len() > k {
			heap.Pop(&hp)
		}
		bfs = bfs[n:]
		n = len(bfs)
	}
	if hp.Len() < k {
		return -1
	}
	return hp.Top()
}
