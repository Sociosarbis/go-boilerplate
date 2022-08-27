package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	bfs := []*TreeNode{root}
	span := make([]int, 0)
	ret := 1
	for len(bfs) != 0 {
		n := len(bfs)
		nextSpan := []int{}
		fmt.Println(bfs)
		for i := 0; i < n; i++ {
			if bfs[i].Left != nil {
				bfs = append(bfs, bfs[i].Left)
				if bfs[i].Right != nil {
					bfs = append(bfs, bfs[i].Right)
					nextSpan = append(nextSpan, 0)
					if i < len(span) {
						nextSpan = append(nextSpan, span[i]*2)
					}
				} else {
					if i < len(span) {
						nextSpan = append(nextSpan, span[i]*2+1)
					}
				}
			} else {
				if bfs[i].Right == nil {
					if i < len(span) {
						if len(nextSpan) != 0 {
							nextSpan[len(nextSpan)-1] += 2 + span[i]*2
						}
					}
				} else {
					bfs = append(bfs, bfs[i].Right)
					if len(nextSpan) != 0 {
						nextSpan[len(nextSpan)-1] += 1
					}
					if i < len(span) {
						nextSpan = append(nextSpan, span[i]*2)
					}
				}
			}
		}
		bfs = bfs[n:]
		if len(bfs) != 0 {
			span = nextSpan[:len(bfs)-1]
			temp := len(bfs)
			for _, width := range span {
				temp += width
			}
			if temp > ret {
				ret = temp
			}
		}
	}
	return ret
}
