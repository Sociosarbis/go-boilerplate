package post

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	preNumToIndex := make(map[int]int, n)
	postNumToIndex := make(map[int]int, n)
	for i := 0; i < n; i++ {
		preNumToIndex[preorder[i]] = i
		postNumToIndex[postorder[i]] = i
	}
	return dfs(preorder, postorder, preNumToIndex, postNumToIndex, 0, n-1)
}

func dfs(preorder, postorder []int, preNumToIndex, postNumToIndex map[int]int, left, right int) *TreeNode {
	num := preorder[left]
	node := &TreeNode{
		Val: num,
	}
	if left+1 <= right {
		r := right
		rightIndex := postNumToIndex[num] - 1
		if rightIndex >= 0 && preorder[left+1] != postorder[rightIndex] {
			r = preNumToIndex[postorder[rightIndex]] - 1
		}
		if left+1 <= r {
			node.Left = dfs(preorder, postorder, preNumToIndex, postNumToIndex, left+1, r)
			left = r
		}
	}
	if left+1 <= right {
		node.Right = dfs(preorder, postorder, preNumToIndex, postNumToIndex, left+1, right)
	}
	return node
}
