package diff

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(node *TreeNode, max int, min int) int {
	if node == nil {
		return 0
	}
	var ret int
	var temp int
	if node.Val <= max {
		ret = max - node.Val
	} else {
		max = node.Val
	}

	if node.Val < min {
		min = node.Val
	} else {
		temp = node.Val - min
		if temp > ret {
			ret = temp
		}
	}

	temp = dfs(node.Left, max, min)

	if temp > ret {
		ret = temp
	}

	temp = dfs(node.Right, max, min)

	if temp > ret {
		ret = temp
	}

	return ret
}
