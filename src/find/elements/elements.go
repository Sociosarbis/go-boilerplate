package elements

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type empty struct{}

type FindElements struct {
	m map[int]empty
}

func Constructor(root *TreeNode) FindElements {
	m := map[int]empty{}
	dfs(root, 0, m)
	return FindElements{
		m,
	}
}

func dfs(node *TreeNode, val int, m map[int]empty) {
	if node == nil {
		return
	}
	m[val] = empty{}
	temp := val * 2
	dfs(node.Left, temp+1, m)
	dfs(node.Right, temp+2, m)
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.m[target]
	return ok
}
