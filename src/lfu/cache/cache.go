package cache

// 不难，但比较麻烦
// https://yuminlee2.medium.com/leetcode-460-lfu-cache-a974db16f24a
type LFUCache struct {
	capacity     int
	size         int
	minFreq      int
	cache        map[int]*LinkedListNode
	frequencyMap map[int]*LinkedList
}

func Constructor(capacity int) LFUCache {
	cache := make(map[int]*LinkedListNode)
	frequencyMap := make(map[int]*LinkedList)
	return LFUCache{
		capacity:     capacity,
		size:         0,
		minFreq:      0,
		cache:        cache,
		frequencyMap: frequencyMap,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, found := this.cache[key]; !found {
		return -1
	}

	tempNode := this.cache[key]
	this.frequencyMap[tempNode.freq].removeNode(tempNode)
	if this.frequencyMap[tempNode.freq].head == nil {
		delete(this.frequencyMap, tempNode.freq)
		if this.minFreq == tempNode.freq {
			this.minFreq += 1
		}
	}

	this.cache[key].freq += 1
	this.initLinkedList(this.cache[key].freq)
	this.frequencyMap[this.cache[key].freq].insertAtTail(this.cache[key])
	return this.cache[key].value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if _, found := this.cache[key]; found {
		this.Get(key)
		this.cache[key].value = value
		return
	}

	if this.size == this.capacity {
		delete(this.cache, this.frequencyMap[this.minFreq].head.key)
		this.frequencyMap[this.minFreq].removeHead()
		if this.frequencyMap[this.minFreq].head == nil {
			delete(this.frequencyMap, this.minFreq)
		}
		this.size -= 1
	}

	this.minFreq = 1
	this.cache[key] = &LinkedListNode{
		key:   key,
		value: value,
		freq:  this.minFreq,
	}

	this.initLinkedList(this.cache[key].freq)
	this.frequencyMap[this.cache[key].freq].insertAtTail(this.cache[key])
	this.size += 1
}

func (this *LFUCache) initLinkedList(freq int) {
	if _, found := this.frequencyMap[freq]; !found {
		this.frequencyMap[freq] = &LinkedList{}
	}
}

type LinkedListNode struct {
	key   int
	value int
	freq  int
	next  *LinkedListNode
	prev  *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func (ll *LinkedList) insertAtTail(node *LinkedListNode) {
	node.prev, node.next = nil, nil
	if ll.tail == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		node.prev = ll.tail
		ll.tail = node
	}
}

func (ll *LinkedList) removeNode(node *LinkedListNode) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == ll.head {
		ll.head = ll.head.next
	}
	if node == ll.tail {
		ll.tail = ll.tail.prev
	}
}

func (ll *LinkedList) removeHead() {
	ll.removeNode(ll.head)
}
