package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val < val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			root.Right = insertIntoMaxTree(root.Right, val)
		}
	}
	return root
}
