package product2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sum(node *TreeNode, mirroredNode *TreeNode) int {
	mirroredNode.Val = node.Val
	if node.Left != nil {
		if mirroredNode.Left == nil {
			mirroredNode.Left = &TreeNode{}
		}
		mirroredNode.Val += sum(node.Left, mirroredNode.Left)
	}
	if node.Right != nil {
		if mirroredNode.Right == nil {
			mirroredNode.Right = &TreeNode{}
		}
		mirroredNode.Val += sum(node.Right, mirroredNode.Right)
	}
	return mirroredNode.Val
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func dfs(node *TreeNode, total int, out *int64) int {
	if node == nil {
		return 0
	}
	if node.Left != nil {
		*out = max(*out, int64(total-node.Left.Val)*int64(node.Left.Val))
		dfs(node.Left, total, out)
	}
	if node.Right != nil {
		*out = max(*out, int64(total-node.Right.Val)*int64(node.Right.Val))
		dfs(node.Right, total, out)
	}
	return int(*out % (1e9 + 7))
}

func maxProduct(root *TreeNode) int {
	var out int64
	var mirroredNode TreeNode
	total := sum(root, &mirroredNode)
	return dfs(&mirroredNode, total, &out)
}
