package node

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 设A长m， B长n，不相等的前缀A长为a, 前缀B长为b
	// 假如A和B相交则，存在相同的片段长l，使 a + l = m, b + l = n
	// 所以当指针A跳到B头，指针B跳到A头，他们碰面的情况下，他们走过的路程分别为a + l + b 和 b + l + a，是相等的，
	// 并且会在A 和 B的相交处碰面
	curA := headA
	curB := headB

	hasSwitchedA := false
	hasSwitchedB := false

	for curA != curB {
		if curA == nil {
			if !hasSwitchedA {
				curA = headB
				hasSwitchedA = true
			}
		} else {
			curA = curA.Next
		}

		if curB == nil {
			if !hasSwitchedB {
				curB = headA
				hasSwitchedB = true
			}
		} else {
			curB = curB.Next
		}

	}
	return curA
}
