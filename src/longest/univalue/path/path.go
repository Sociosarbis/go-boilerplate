package path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, prevVal int) int {
	if node == nil {
		return 0
	}
	if prevVal > 1000 || node.Val == prevVal {
		if prevVal > 1000 {
			return 1 + dfs(node.Left, node.Val) + dfs(node.Right, node.Val)
		} else {
			leftCount := dfs(node.Left, node.Val)
			rightCount := dfs(node.Right, node.Val)
			if leftCount > rightCount {
				return 1 + leftCount
			} else {
				return 1 + rightCount
			}
		}
	}
	return 0
}

func compute(node *TreeNode, max *int) {
	ret := dfs(node, 1001)
	if ret > *max {
		*max = ret
	}
	if node.Left != nil {
		compute(node.Left, max)
	}
	if node.Right != nil {
		compute(node.Right, max)
	}
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 1
	compute(root, &max)
	return max - 1
}
