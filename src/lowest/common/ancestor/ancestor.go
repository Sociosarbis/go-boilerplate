package ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	left := findNum(root, []*TreeNode{}, p.Val)
	right := findNum(root, []*TreeNode{}, q.Val)
	i := 0
	for i+1 < len(left) && i+1 < len(right) && left[i+1].Val == right[i+1].Val {
		i++
	}
	return left[i]
}

func findNum(node *TreeNode, path []*TreeNode, targetValue int) []*TreeNode {
	if node == nil {
		return nil
	}
	path = append(path, node)
	if node.Val == targetValue {
		return path
	}

	res := findNum(node.Left, path, targetValue)
	if res != nil {
		return res
	}
	return findNum(node.Right, path, targetValue)
}
