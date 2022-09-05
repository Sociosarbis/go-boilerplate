package subtrees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	node  *TreeNode
	count int
}

func dfs(node *TreeNode, m map[string]Item) string {
	if node == nil {
		return "-"
	}
	var s string
	if node.Left == nil && node.Right == nil {
		s = fmt.Sprintf("%v", node.Val)
	} else {
		s = dfs(node.Left, m) + " " + dfs(node.Right, m) + " " + fmt.Sprintf("%v", node.Val)
	}
	if item, ok := m[s]; ok {
		item.count++
		m[s] = item
	} else {
		m[s] = Item{
			node:  node,
			count: 1,
		}
	}
	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := map[string]Item{}
	dfs(root, m)
	ret := []*TreeNode{}
	for _, item := range m {
		if item.count > 1 {
			ret = append(ret, item.node)
		}
	}
	return ret
}
