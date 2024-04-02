package fbt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	return dfs(n, make([][]*TreeNode, n+1))
}

func dfs(n int, dp [][]*TreeNode) []*TreeNode {
	if n%2 != 1 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{{}}
	}
	if len(dp[n]) != 0 {
		return dp[n]
	}
	n--
	ret := []*TreeNode{}
	for i := 1; i < n; i += 2 {
		left := dfs(i, dp)
		right := dfs(n-i, dp)
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				ret = append(ret, &TreeNode{
					Left:  left[j],
					Right: right[k],
				})
			}
		}
	}
	dp[n+1] = ret
	return ret
}
