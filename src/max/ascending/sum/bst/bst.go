package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	max   int
	min   int
	sum   int
	isBST bool
}

func maxSumBST(root *TreeNode) int {
	ret := 0
	dfs(root, &ret)
	return ret
}

func dfs(node *TreeNode, res *int) Info {
	var ret = Info{
		max:   node.Val,
		min:   node.Val,
		sum:   node.Val,
		isBST: true,
	}
	if node.Left != nil {
		childInfo := dfs(node.Left, res)
		if !childInfo.isBST || childInfo.max >= node.Val {
			ret.isBST = false
		} else {
			ret.min = childInfo.min
			ret.sum += childInfo.sum
		}
	}

	if node.Right != nil {
		childInfo := dfs(node.Right, res)
		if !childInfo.isBST || childInfo.min <= node.Val {
			ret.isBST = false
		} else {
			ret.max = childInfo.max
			ret.sum += childInfo.sum
		}
	}
	if ret.isBST && ret.sum > *res {
		*res = ret.sum
	}
	return ret
}
