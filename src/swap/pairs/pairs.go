package pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var ret *ListNode
	var tail *ListNode

	for head != nil {
		if ret == nil {
			if head.Next == nil {
				ret = head
				head = nil
			} else {
				ret = head.Next
				temp := head.Next.Next
				ret.Next = head
				tail = head
				tail.Next = nil
				head = temp
			}
		} else {
			if head.Next == nil {
				tail.Next = head
				head = nil
			} else {
				tail.Next = head.Next
				temp := head.Next.Next
				tail.Next.Next = head
				tail = head
				tail.Next = nil
				head = temp
			}
		}
	}
	return ret
}
