package coins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	return dfs(root, nil)
}

func dfs(child *TreeNode, parent *TreeNode) int {
	if child != nil {
		ret := 0
		ret += dfs(child.Left, child)
		ret += dfs(child.Right, child)
		if child.Val < 1 {
			ret += 1 - child.Val
			parent.Val -= 1 - child.Val
		} else if child.Val > 1 {
			ret += child.Val - 1
			parent.Val += child.Val - 1
		}
		return ret
	}
	return 0
}
