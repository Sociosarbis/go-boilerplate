package leaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, temp int, depth int) int {
	if node == nil {
		return 0
	}
	ret := 0
	if node.Val == 1 {
		temp |= 1 << depth
	}
	if node.Left == nil && node.Right == nil {
		val := 0
		for i := 0; i <= depth; i++ {
			val |= (temp & 1) << (depth - i)
			temp >>= 1
		}
		ret += val
	} else {
		ret += dfs(node.Left, temp, depth+1)
		ret += dfs(node.Right, temp, depth+1)
	}
	return ret
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0, 0)
}
