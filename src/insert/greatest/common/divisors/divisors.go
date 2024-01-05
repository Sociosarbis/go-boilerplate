package divisors

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur.Next != nil {
		oldNext := cur.Next
		cur, cur.Next = oldNext, &ListNode{
			Val:  gcd(cur.Val, oldNext.Val),
			Next: oldNext,
		}
	}
	return head
}
