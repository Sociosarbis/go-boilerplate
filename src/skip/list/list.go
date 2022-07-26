package list

import "math/rand"

type Node struct {
	val  int
	next *Node
	down *Node
}

type Skiplist struct {
	head *Node
}

func Constructor() Skiplist {
	return Skiplist{
		head: &Node{
			-1,
			nil,
			nil,
		},
	}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for cur != nil {
		for cur.next != nil && cur.next.val < target {
			cur = cur.next
		}
		if cur.next != nil && cur.next.val == target {
			return true
		}
		cur = cur.down
	}
	return false
}

func (this *Skiplist) Add(num int) {
	st := []*Node{}
	cur := this.head
	for cur != nil {
		for cur.next != nil && cur.next.val < num {
			cur = cur.next
		}
		st = append(st, cur)
		cur = cur.down
	}
	insert := true
	var down *Node
	for insert && len(st) != 0 {
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		cur.next = &Node{
			num,
			cur.next,
			down,
		}
		down = cur.next
		insert = rand.Intn(2) == 0
	}
	if insert {
		this.head = &Node{
			-1,
			nil,
			this.head,
		}
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur := this.head
	found := false
	for cur != nil {
		for cur.next != nil && cur.next.val < num {
			cur = cur.next
		}
		if cur.next != nil && cur.next.val == num {
			found = true
			cur.next = cur.next.next
		}
		cur = cur.down
	}
	return found
}
