package nodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	foward := removeNodes(head.Next)
	if foward == nil || head.Val >= foward.Val {
		head.Next = foward
		return head
	} else {
		return foward
	}
}
