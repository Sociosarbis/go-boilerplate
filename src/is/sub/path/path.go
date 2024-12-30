package path

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	return dfs2(root, head)
}

func dfs1(node *TreeNode, other *ListNode) bool {
	if other == nil {
		return true
	}
	if node == nil {
		return false
	}
	if node.Val == other.Val {
		res := dfs1(node.Left, other.Next)
		if res {
			return res
		}
		return dfs1(node.Right, other.Next)
	}
	return false
}

func dfs2(node *TreeNode, head *ListNode) bool {
	if node == nil {
		return false
	}
	if dfs1(node, head) {
		return true
	}
	res := dfs2(node.Left, head)
	if res {
		return res
	}
	return dfs2(node.Right, head)
}
