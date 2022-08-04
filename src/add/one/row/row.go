package row

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, val int, depth int, curDepth int) {
	if curDepth+1 == depth {
		if node.Left != nil {
			node.Left = &TreeNode{
				val,
				node.Left,
				nil,
			}
		} else {
			node.Left = &TreeNode{
				val,
				nil,
				nil,
			}
		}

		if node.Right != nil {
			node.Right = &TreeNode{
				val,
				nil,
				node.Right,
			}
		} else {
			node.Right = &TreeNode{
				val,
				nil,
				nil,
			}
		}
	} else {
		if node.Left != nil {
			dfs(node.Left, val, depth, curDepth+1)
		}
		if node.Right != nil {
			dfs(node.Right, val, depth, curDepth+1)
		}
	}
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			val,
			root,
			nil,
		}
	} else {
		dfs(root, val, depth, 1)
		return root
	}
}
