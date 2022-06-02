package node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var parent *TreeNode
	cur := root
	for cur != nil && cur.Val != key {
		parent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if cur != nil {
		var replaceNode *TreeNode
		if cur.Right != nil {
			replaceNode = cur.Right
		} else if cur.Left != nil {
			replaceNode = cur.Left
		} else {
			replaceNode = nil
		}
		if parent == nil {
			root = replaceNode
		} else {
			if parent.Left == cur {
				parent.Left = replaceNode
			} else {
				parent.Right = replaceNode
			}
		}
		leftChild := cur.Left
		if cur.Right != nil {
			cur = cur.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = leftChild
		}
	}
	return root
}
