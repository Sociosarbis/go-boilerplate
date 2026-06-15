package middle

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var n int
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	if n == 1 {
		return nil
	}
	cur = head
	for i := n/2 - 1; i > 0; i-- {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	} else {
		cur.Next = nil
	}
	return head
}
