package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	valToNode := make(map[int]*TreeNode, len(descriptions)*2)
	for _, description := range descriptions {
		parent, child, isLeft := description[0], description[1], description[2] == 1
		var parentNode *TreeNode
		if node, ok := valToNode[parent]; ok {
			parentNode = node
		} else {
			parentNode = newTreeNode(parent)
			valToNode[parent] = parentNode
		}
		var childNode *TreeNode
		if node, ok := valToNode[child]; ok {
			childNode = node
		} else {
			childNode = newTreeNode(child)
			valToNode[child] = childNode
		}
		if isLeft {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}

	for _, description := range descriptions {
		delete(valToNode, description[1])
	}
	var out *TreeNode
	for _, value := range valToNode {
		out = value
		break
	}
	return out
}
