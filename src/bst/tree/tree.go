package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(node *TreeNode, acc int) int {
	if node == nil {
		return 0
	}
	if node.Right != nil {
		acc = dfs(node.Right, acc)
	}
	node.Val += acc
	if node.Left != nil {
		acc = dfs(node.Left, node.Val)
	} else {
		acc = node.Val
	}
	return acc
}
