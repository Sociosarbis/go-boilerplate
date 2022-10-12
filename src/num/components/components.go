package components

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	m := make([]bool, n)
	for _, num := range nums {
		m[num] = true
	}
	ret := 0
	cur = head
	hasFirst := false
	for cur != nil {
		if m[cur.Val] {
			if !hasFirst {
				hasFirst = true
				ret++
			}
		} else {
			hasFirst = false
		}
	}
	return ret
}
