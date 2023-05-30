package nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := make(map[int]struct{}, len(to_delete))
	for _, v := range to_delete {
		m[v] = struct{}{}
	}
	temp := TreeNode{
		Left: root,
	}
	return dfs(&temp, temp.Left, m)
}

func dfs(p *TreeNode, c *TreeNode, to_delete map[int]struct{}) []*TreeNode {
	ret := []*TreeNode{}
	if _, ok := to_delete[c.Val]; ok {
		if c.Left != nil {
			ret = append(ret, dfs(c, c.Left, to_delete)...)
		}
		if c.Right != nil {
			ret = append(ret, dfs(c, c.Right, to_delete)...)
		}
		if p.Left == c {
			p.Left = nil
		} else {
			p.Right = nil
		}
		delete(to_delete, c.Val)
	} else {
		ret = append(ret, c)
		if c.Left != nil {
			temp := dfs(c, c.Left, to_delete)
			if len(temp) != 0 {
				if c.Left != nil {
					ret = append(ret, temp[1:]...)
				} else {
					ret = append(ret, temp...)
				}
			}
		}
		if c.Right != nil {
			temp := dfs(c, c.Right, to_delete)
			if len(temp) != 0 {
				if c.Right != nil {
					ret = append(ret, temp[1:]...)
				} else {
					ret = append(ret, temp...)
				}
			}
		}
	}
	return ret
}
