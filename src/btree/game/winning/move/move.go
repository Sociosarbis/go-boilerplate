package move

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(node *TreeNode) int {
	if node != nil {
		return 1 + countNodes(node.Left) + countNodes(node.Right)
	}
	return 0
}

func findNode(node *TreeNode, x int) *TreeNode {
	if node != nil {
		if node.Val == x {
			return node
		}
		ret := findNode(node.Left, x)
		if ret != nil {
			return ret
		}
		return findNode(node.Right, x)
	}
	return nil
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	node := findNode(root, x)
	leftCount := countNodes(node.Left)
	rightCount := countNodes(node.Right)
	otherCount := n - leftCount - rightCount - 1
	if leftCount > rightCount+otherCount+1 || rightCount > leftCount+otherCount+1 || otherCount > leftCount+rightCount+1 {
		return true
	}
	return false
}
