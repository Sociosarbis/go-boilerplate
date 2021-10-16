package smallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ret, _, _ := dfs(root, 1, k)
	return ret
}

func dfs(node *TreeNode, i int, k int) (int, int, bool) {
	if node.Left != nil {
		v, index, ok := dfs(node.Left, i, k)
		if ok {
			return v, index, ok
		}
		i = index
	}
	if i == k {
		return node.Val, i + 1, true
	}
	value := node.Val
	i += 1
	if node.Right != nil {
		v, index, ok := dfs(node.Right, i, k)
		if ok {
			return v, index, ok
		}
		i = index
		value = v
	}
	return value, i, false
}
