package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, i, j int, mat *[][]string) {
	(*mat)[i][j] = fmt.Sprint(node.Val)

	if node.Left != nil {
		dfs(node.Left, i+1, j-(1<<(len(*mat)-2-i)), mat)
	}
	if node.Right != nil {
		dfs(node.Right, i+1, j+(1<<(len(*mat)-2-i)), mat)
	}
}

func printTree(root *TreeNode) [][]string {
	bfs := []*TreeNode{root}
	height := 0
	for len(bfs) != 0 {
		n := len(bfs)
		height++
		for i := 0; i < n; i++ {
			if bfs[i].Left != nil {
				bfs = append(bfs, bfs[i].Left)
			}
			if bfs[i].Right != nil {
				bfs = append(bfs, bfs[i].Right)
			}
		}
		bfs = bfs[n:]
	}
	m := height
	n := (1 << height) - 1
	ret := make([][]string, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]string, n)
	}
	dfs(root, 0, (n-1)>>1, &ret)
	return ret
}
