package nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, -10001)
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	ret := 0
	if node.Val >= val {
		val = node.Val
		ret += 1
	}
	return ret + dfs(node.Right, val) + dfs(node.Left, val)
}
