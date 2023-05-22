package subset

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	dfs(root, 0, limit)
	if root.Left == nil && root.Right == nil && root.Val < limit {
		return nil
	}
	return root
}

func dfs(node *TreeNode, temp int, limit int) bool {
	val := temp + node.Val
	if node.Left == nil && node.Right == nil {
		return val >= limit
	}

	if node.Left != nil {
		res := dfs(node.Left, val, limit)
		if !res {
			node.Left = nil
		}
	}

	if node.Right != nil {
		res := dfs(node.Right, val, limit)
		if !res {
			node.Right = nil
		}
	}

	return node.Left != nil || node.Right != nil
}
