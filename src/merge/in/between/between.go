package between

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	i := 0
	head := list1
	for i+1 < a {
		head = head.Next
		i++
	}
	tail := head
	for i < b+1 {
		tail = tail.Next
		i++
	}
	tail2 := list2
	for tail2.Next != nil {
		tail2 = tail2.Next
	}
	tail2.Next = tail

	head.Next = list2
	return list1
}
