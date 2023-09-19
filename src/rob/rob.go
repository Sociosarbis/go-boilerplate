package rob

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := dfs(root)
	if res[0] > res[1] {
		return res[0]
	}
	return res[1]
}

func dfs(node *TreeNode) []int {
	ret := []int{0, node.Val}

	if node.Left != nil {
		res := dfs(node.Left)
		ret[1] += res[0]
		if res[0] > res[1] {
			ret[0] += res[0]
		} else {
			ret[0] += res[1]
		}
	}

	if node.Right != nil {
		res := dfs(node.Right)
		ret[1] += res[0]
		if res[0] > res[1] {
			ret[0] += res[0]
		} else {
			ret[0] += res[1]
		}
	}
	return ret
}
