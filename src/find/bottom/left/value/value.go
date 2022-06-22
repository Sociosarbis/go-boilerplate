package value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, level int) (int, int) {
	if node.Left == nil && node.Right == nil {
		return node.Val, level
	}

	val := node.Val
	lv := level

	if node.Left != nil {
		nextVal, nextLevel := dfs(node.Left, level+1)
		lv = nextLevel
		val = nextVal
	}

	if node.Right != nil {
		nextVal, nextLevel := dfs(node.Right, level+1)
		if nextLevel > lv {
			lv = nextLevel
			val = nextVal
		}
	}
	return val, lv
}

func findBottomLeftValue(root *TreeNode) int {
	ret, _ := dfs(root, 0)
	return ret
}
