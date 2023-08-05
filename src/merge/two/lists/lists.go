package lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ret *ListNode

	var tail *ListNode

	for list1 != nil || list2 != nil {
		if list1 == nil || (list2 != nil && list1.Val > list2.Val) {
			if ret == nil {
				ret = list2
				tail = ret
			} else {
				tail.Next = list2
				tail = tail.Next
			}
			list2 = list2.Next
		} else {
			if ret == nil {
				ret = list1
				tail = ret
			} else {
				tail.Next = list1
				tail = tail.Next
			}
			list1 = list1.Next
		}
	}
	return ret
}
