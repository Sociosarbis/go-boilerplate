package list

type ListNode struct {
	Val  int
	Next *ListNode
}

type empty struct{}

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]empty, len(nums))
	for _, num := range nums {
		m[num] = empty{}
	}
	cur := head
	var prev *ListNode
	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			if cur == head {
				head = cur.Next
			} else {
				prev.Next = cur.Next
			}
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return head
}
