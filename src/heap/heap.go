package heap

import "container/heap"

type Heap[T any] struct {
	items []T
	less  func(a, b T) bool
}

func New[T any](items []T, less func(a, b T) bool) Heap[T] {
	h := Heap[T]{items, less}
	heap.Init(&h)
	return h
}

func (h Heap[T]) Len() int           { return len(h.items) }
func (h Heap[T]) Less(i, j int) bool { return h.less(h.items[i], h.items[j]) }
func (h Heap[T]) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *Heap[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

func (h *Heap[T]) Pop() any {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[0 : n-1]
	return x
}

func (h *Heap[T]) Top() T {
	return h.items[0]
}
