package levels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	layer := []*TreeNode{root}
	index := 0
	for len(layer) != 0 {
		index++
		n := len(layer)
		for i := 0; i < n; i++ {
			if layer[i].Left != nil {
				layer = append(layer, layer[i].Left, layer[i].Right)
			}
		}
		layer = layer[n:]
		if index&1 == 1 {
			l := 0
			r := len(layer) - 1
			for l < r {
				layer[l].Val, layer[r].Val = layer[r].Val, layer[l].Val
				l++
				r--
			}
		}
	}
	return root
}
