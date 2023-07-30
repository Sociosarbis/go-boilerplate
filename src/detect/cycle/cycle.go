package cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	if fast != nil {
		fast = head.Next
	}
	for fast != nil && fast != slow {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	if fast == nil {
		return nil
	}
	// 这时头到相遇点的距离是环路程的倍数
	// 所以分别从头和相遇点以同样的速度行走，一定会在环起始点相遇
	slow = head
	fast = fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
