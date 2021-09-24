package flatten

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	ret, _ := flattenDFS(root)
	return ret
}

func flattenDFS(root *Node) (*Node, *Node) {
	head := root
	tail := root
	if head != nil {
		cur := head
		for cur != nil {
			next := cur.Next
			if cur.Child != nil {
				subHead, subTail := flattenDFS(cur.Child)
				cur.Next = subHead
				cur.Child = nil
				subHead.Prev = cur
				subTail.Next = next
				if next != nil {
					next.Prev = subTail
				}
				cur = subTail
			}
			if next == nil {
				tail = cur
			}
			cur = next
		}
	}
	return head, tail
}
