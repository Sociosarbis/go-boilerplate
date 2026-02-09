package bst

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func collect(node *TreeNode, out *[]int) {
	if node != nil {
		*out = append(*out, node.Val)
		collect(node.Left, out)
		collect(node.Right, out)
	}
}

func build(out []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return &TreeNode{
		Val:   out[mid],
		Left:  build(out, l, mid-1),
		Right: build(out, mid+1, r),
	}
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := []int{}
	collect(root, &nums)
	sort.Ints(nums)
	return build(nums, 0, len(nums)-1)
}
