package value

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var out int
	for head != nil {
		out = (out << 1) | head.Val
		head = head.Next
	}
	return out
}
