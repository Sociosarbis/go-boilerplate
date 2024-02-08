package cousins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	left := dfs(root, 0, 0, x)
	right := dfs(root, 0, 0, y)
	return left != nil && right != nil && left[0] != right[0] && left[1] == right[1]
}

func dfs(node *TreeNode, parent int, level int, x int) *[2]int {
	if node == nil {
		return nil
	}
	if node.Val == x {
		return &[2]int{parent, level}
	}
	for _, child := range [2]*TreeNode{node.Left, node.Right} {
		res := dfs(child, node.Val, level+1, x)
		if res != nil {
			return res
		}
	}
	return nil
}
