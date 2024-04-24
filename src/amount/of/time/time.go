package time

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	ret := 0
	dfs(root, start, 0, &ret)
	if ret <= 1 {
		return 0
	}
	return ret - 1
}

func dfs(node *TreeNode, start int, temp int, out *int) int {
	if node == nil {
		return 0
	}

	if temp == 0 {
		if node.Val == start {
			dfs(node.Left, start, 1, out)
			dfs(node.Right, start, 1, out)
			if *out < 1 {
				*out = 1
			}
			return 1
		} else {
			temp := dfs(node.Left, start, 0, out)
			if temp != 0 {
				dfs(node.Right, start, temp+1, out)
				if *out < temp+1 {
					*out = temp + 1
				}
				return temp + 1
			}
			temp = dfs(node.Right, start, 0, out)
			if temp != 0 {
				dfs(node.Left, start, temp+1, out)
				if *out < temp+1 {
					*out = temp + 1
				}
				return temp + 1
			}
			return 0
		}
	} else {
		dfs(node.Left, start, temp+1, out)
		dfs(node.Right, start, temp+1, out)
		if *out < temp+1 {
			*out = temp + 1
		}
		return temp + 1
	}
}
