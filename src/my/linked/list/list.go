package list

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	head *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	n := this.getNode(index)
	if n != nil {
		return n.val
	}
	return -1
}

func (this *MyLinkedList) getNode(index int) *node {
	if this.head == nil {
		return nil
	}
	cur := this.head
	for i := 0; i < index; i++ {
		if cur != nil {
			cur = cur.next
		} else {
			return nil
		}
	}
	if cur != nil {
		return cur
	} else {
		return nil
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := node{
		val,
		nil,
	}
	n.next = this.head
	this.head = &n
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
	} else {
		cur := this.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = &node{
			val,
			nil,
		}
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else {
		p := this.getNode(index - 1)
		if p != nil {
			next := p.next
			n := node{
				val,
				next,
			}
			p.next = &n
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.head != nil {
			this.head = this.head.next
		}
	} else {
		p := this.getNode(index - 1)
		c := p.next
		if c == nil {
			p.next = nil
		} else {
			p.next = c.next
		}
	}
}
