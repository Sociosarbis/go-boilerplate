package traversal

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Val int
	X   int
	Y   int
}

func verticalTraversal(root *TreeNode) [][]int {
	items := []Item{}
	dfs(root, 0, 0, &items)
	sort.Slice(items, func(i, j int) bool {
		if items[i].X < items[j].X ||
			items[i].X == items[j].X &&
				(items[i].Y < items[j].Y ||
					(items[i].Y == items[j].Y && items[i].Val < items[j].Val)) {
			return true
		}
		return false
	})
	ret := [][]int{{items[0].Val}}
	x := items[0].X
	for i := 1; i < len(items); i += 1 {
		if items[i].X != x {
			ret = append(ret, []int{})
			x = items[i].X
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], items[i].Val)
	}
	return ret
}

func dfs(node *TreeNode, x, y int, items *[]Item) {
	if node != nil {
		*items = append(*items, Item{
			node.Val,
			x,
			y,
		})
		dfs(node.Left, x-1, y+1, items)
		dfs(node.Right, x+1, y+1, items)
	}
}
