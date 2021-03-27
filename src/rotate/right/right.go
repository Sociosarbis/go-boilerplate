package right

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	i := k % count
	if i != 0 {
		cur = head
		for j := 0; j < count-i-1; j++ {
			cur = cur.Next
		}
		prev := cur
		for cur.Next != nil {
			cur = cur.Next
		}
		prevHead := head
		head = prev.Next
		prev.Next = nil
		cur.Next = prevHead
	}
	return head
}
