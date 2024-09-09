package nodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	cur := head
	for {
		first := cur
		for cur.Next.Val != 0 {
			first.Val += cur.Next.Val
			cur = cur.Next
		}
		// cur无下个节点，或者cur的下个节点是0
		for cur.Next != nil && cur.Next.Val == 0 {
			cur = cur.Next
		}
		// cur无下个节点，或者cur的下个节点不是0
		if cur.Next == nil {
			first.Next = nil
			break
		} else {
			first.Next = cur
		}
	}
	return head
}
