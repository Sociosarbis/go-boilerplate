package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
