package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, target int) bool {
	if node == nil {
		return true
	}

	return node.Val == target && dfs(node.Left, target) && dfs(node.Right, target)
}

func isUnivalTree(root *TreeNode) bool {
	return dfs(root, root.Val)
}
