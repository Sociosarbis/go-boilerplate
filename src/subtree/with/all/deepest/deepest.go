package deepest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, level int) (*TreeNode, int) {
	out := node
	maxLevel := level
	if node.Left != nil {
		out, maxLevel = dfs(node.Left, level+1)
	}
	if node.Right != nil {
		n, l := dfs(node.Right, level+1)
		if l > maxLevel {
			return n, l
		} else if l == maxLevel {
			if out == nil {
				return n, l
			} else {
				return node, l
			}
		}
	}
	return out, maxLevel
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	out, _ := dfs(root, 0)
	return out
}
