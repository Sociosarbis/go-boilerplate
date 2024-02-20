package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	numToIdx := make(map[int]int, n)
	for i, num := range inorder {
		numToIdx[num] = i
	}
	ret, _ := dfs(preorder, numToIdx, 0, 0, n-1)
	return ret
}

func dfs(preorder []int, numToIdx map[int]int, i, left, right int) (*TreeNode, int) {
	num := preorder[i]
	index := numToIdx[num]
	ret := &TreeNode{
		Val: num,
	}
	if index-1 >= left {
		left, nextI := dfs(preorder, numToIdx, i+1, left, index-1)
		ret.Left = left
		i = nextI
	}

	if index+1 <= right {
		right, nextI := dfs(preorder, numToIdx, i+1, index+1, right)
		ret.Right = right
		i = nextI
	}
	return ret, i
}
