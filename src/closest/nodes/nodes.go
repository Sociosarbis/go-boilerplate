package nodes

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	ret := make([][]int, len(queries))
	values := []int{}
	dfs(root, &values)
	n := len(values)
	for i, q := range queries {
		index := sort.SearchInts(values, q)
		ret[i] = make([]int, 2)
		if index >= n {
			ret[i][0] = values[index-1]
			ret[i][1] = -1
		} else {
			ret[i][1] = values[index]
			if values[index] == q {
				ret[i][0] = q
			} else {
				if index > 0 {
					ret[i][0] = values[index-1]
				} else {
					ret[i][0] = -1
				}
			}
		}
	}
	return ret
}

func dfs(node *TreeNode, out *[]int) {
	if node.Left != nil {
		dfs(node.Left, out)
	}
	*out = append(*out, node.Val)
	if node.Right != nil {
		dfs(node.Right, out)
	}
}
