package lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	queue := make([]*ListNode, 0, len(lists))
	for _, node := range lists {
		if node == nil {
			continue
		}
		l := 0
		r := len(queue) - 1
		for l <= r {
			mid := (l + r) >> 1
			if node.Val < queue[mid].Val {
				l = mid + 1
			} else if node.Val > queue[mid].Val {
				r = mid - 1
			} else {
				l = mid
				break
			}
		}
		if l >= len(queue) {
			queue = append(queue, node)
		} else {
			queue = append(queue, nil)
			copy(queue[l+1:], queue[l:])
			queue[l] = node
		}
	}
	ret := &ListNode{}
	tail := ret
	for len(queue) != 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		tail.Next = node
		tail = tail.Next
		node = node.Next
		if node != nil {
			l := 0
			r := len(queue) - 1
			for l <= r {
				mid := (l + r) >> 1
				if node.Val < queue[mid].Val {
					l = mid + 1
				} else if node.Val > queue[mid].Val {
					r = mid - 1
				} else {
					l = mid
					break
				}
			}
			if l >= len(queue) {
				queue = append(queue, node)
			} else {
				queue = append(queue, nil)
				copy(queue[l+1:], queue[l:])
				queue[l] = node
			}
		}
	}
	return ret.Next
}
