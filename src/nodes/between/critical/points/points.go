package points

type ListNode struct {
	Val  int
	Next *ListNode
}

func isCritical(prev, cur, next int) bool {
	return (cur > prev && cur > next) || (cur < prev && cur < next)
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	prev, cur := head, head.Next
	if cur.Next == nil {
		return []int{-1, -1}
	}
	next := cur.Next
	var i int
	maxPairs := make([]int, 0, 2)
	if isCritical(prev.Val, cur.Val, next.Val) {
		maxPairs = append(maxPairs, i)
	}
	var min int = 1e5
	for next != nil {
		prev, cur, next = cur, next, next.Next
		i++
		if next == nil {
			break
		}
		if isCritical(prev.Val, cur.Val, next.Val) {
			n := len(maxPairs)
			if n > 0 {
				if i-maxPairs[n-1] < min {
					min = i - maxPairs[n-1]
				}
				if n == 1 {
					maxPairs = append(maxPairs, i)
				} else {
					maxPairs[n-1] = i
				}
			} else {
				maxPairs = append(maxPairs, i)
			}
		}
	}
	if len(maxPairs) < 2 {
		return []int{-1, -1}
	}
	return []int{min, maxPairs[1] - maxPairs[0]}
}
