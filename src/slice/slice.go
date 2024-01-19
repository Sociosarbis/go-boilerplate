package slice

type Sortable[T any] struct {
	arr  []T
	less func(a, b T) bool
}

func NewSortable[T any](arr []T, less func(a, b T) bool) Sortable[T] {
	return Sortable[T]{
		arr,
		less,
	}
}

func (this Sortable[T]) Len() int {
	return len(this.arr)
}

func (this Sortable[T]) Swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func (this Sortable[T]) Less(i, j int) bool {
	return this.less(this.arr[i], this.arr[j])
}
