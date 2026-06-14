package sum

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var n int
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	var out int
	halfN := n / 2
	stack := make([]int, 0, halfN)
	for i := 0; i < halfN; i++ {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := halfN - 1; i >= 0; i-- {
		if head.Val+stack[i] > out {
			out = head.Val + stack[i]
		}
		head = head.Next
	}
	return out
}
