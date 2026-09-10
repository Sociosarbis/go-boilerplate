package substree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, out *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	sum, count := node.Val, 1
	s, c := dfs(node.Left, out)
	sum += s
	count += c
	s, c = dfs(node.Right, out)
	sum += s
	count += c
	if node.Val == sum/count {
		*out++
	}
	return sum, count
}

func averageOfSubtree(root *TreeNode) int {
	var out int
	dfs(root, &out)
	return out
}
