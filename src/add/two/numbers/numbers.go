package numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	revL1 := rev(l1)
	revL2 := rev(l2)
	var head *ListNode
	var temp int
	for revL1 != nil || revL2 != nil {
		var node *ListNode
		if revL1 != nil {
			if revL2 != nil {
				revL1.Val += revL2.Val + temp
				revL2 = revL2.Next
			} else {
				revL1.Val += temp
			}
			node = revL1
			revL1 = revL1.Next
		} else {
			revL2.Val += temp
			node = revL2
			revL2 = revL2.Next
		}
		if node.Val >= 10 {
			node.Val -= 10
			temp = 1
		} else {
			temp = 0
		}
		node.Next = head
		head = node
	}
	if temp != 0 {
		head = &ListNode{
			Val:  1,
			Next: head,
		}
	}
	return head
}

func rev(l *ListNode) *ListNode {
	var head *ListNode
	for l != nil {
		next := l.Next
		l.Next = head
		head = l
		l = next
	}
	return head
}
