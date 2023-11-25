package paths

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	counter := make([]int, 10)
	return dfs(root, counter, 0, 0)
}

func dfs(node *TreeNode, counter []int, size int, oddCount int) int {
	size += 1
	counter[node.Val] += 1
	if counter[node.Val]&1 == 1 {
		oddCount += 1
	} else {
		oddCount -= 1
	}
	if node.Left == nil && node.Right == nil {
		counter[node.Val] -= 1
		if size&1 == oddCount {
			return 1
		}
		return 0
	} else {
		ret := 0
		if node.Left != nil {
			ret += dfs(node.Left, counter, size, oddCount)
		}
		if node.Right != nil {
			ret += dfs(node.Right, counter, size, oddCount)
		}
		counter[node.Val] -= 1
		return ret
	}
}
