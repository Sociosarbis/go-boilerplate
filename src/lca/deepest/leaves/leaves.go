package leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Res struct {
	node  *TreeNode
	depth int
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res := dfs(root, 0)
	return res.node
}

func dfs(node *TreeNode, depth int) Res {
	if node == nil {
		return Res{
			nil,
			depth,
		}
	}
	if node.Left == nil && node.Right == nil {
		return Res{
			node,
			depth,
		}
	}

	res1 := dfs(node.Left, depth+1)
	res2 := dfs(node.Right, depth+1)
	if res1.node == nil || res1.depth < res2.depth {
		return res2
	} else if res2.node == nil || res2.depth < res1.depth {
		return res1
	} else {
		return Res{
			node,
			res1.depth,
		}
	}
}
