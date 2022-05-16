package successor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, target int, temp *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val > target {
		ret := dfs(node.Left, target, temp)
		if ret != nil {
			return ret
		} else if temp != nil {
			return node
		}
	} else {
		if node.Val == target {
			*temp = *node
		}
		return dfs(node.Right, target, temp)
	}
	return nil
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var temp TreeNode
	return dfs(root, p.Val, &temp)
}
