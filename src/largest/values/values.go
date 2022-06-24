package values

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	ret := []int{}
	for len(stack) != 0 {
		n := len(stack)
		max := stack[0].Val
		for i := 0; i < n; i++ {
			if stack[i].Val > max {
				max = stack[i].Val
			}

			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}

			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[n:]
		ret = append(ret, max)
	}
	return ret
}
