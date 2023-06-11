package sublists

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	cur := head
	temp := 0
	visited := map[int]*ListNode{}
	for cur != nil {
		if cur.Val != 0 {
			temp += cur.Val
			if temp == 0 {
				visited = map[int]*ListNode{}
				head = cur.Next
			} else {
				if node, ok := visited[temp]; ok {
					tempNode := node.Next
					tempTemp := temp
					for tempNode != cur {
						tempTemp += tempNode.Val
						delete(visited, tempTemp)
						tempNode = tempNode.Next
					}
					node.Next = cur.Next
				} else {
					visited[temp] = cur
				}
			}
		} else {
			if head == cur {
				head = cur.Next
			} else {
				if node, ok := visited[temp]; ok {
					node.Next = cur.Next
				}
			}
		}
		cur = cur.Next
	}
	return head
}
