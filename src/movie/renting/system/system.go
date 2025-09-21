package system

import (
	"container/heap"

	hp "github.com/boilerplate/src/heap"
)

type rented struct {
	included map[int]map[int]int
	queue    hp.Heap[[]int]
}

func newRented(n int) *rented {
	included := make(map[int]map[int]int, n)
	out := &rented{
		included: included,
		queue: hp.New[[]int](nil, func(a []int, b []int) bool {
			return a[2] < b[2] || (a[2] == b[2] && a[0] < b[0]) || (a[2] == b[2] && a[0] == b[0] && a[1] < b[1])
		}),
	}
	return out
}

func add(included map[int]map[int]int, item []int) {
	shop, movie, price := item[0], item[1], item[2]
	if m, ok := included[shop]; ok {
		m[movie] = price
	} else {
		included[shop] = map[int]int{
			movie: price,
		}
	}
}

func (this *rented) add(item []int) {
	add(this.included, item)
	heap.Push(&this.queue, item)
}

func (this *rented) remove(shop, movie int) int {
	price := this.included[shop][movie]
	delete(this.included[shop], movie)
	return price
}

func has(included map[int]map[int]int, shop, movie int) bool {
	if m, ok := included[shop]; ok {
		if _, ok := m[movie]; ok {
			return true
		}
	}
	return false
}

func (this *rented) getTop5() [][]int {
	out := make([][]int, 0, 5)
	prices := make([]int, 0, 5)
	for len(out) < 5 && this.queue.Len() > 0 {
		top := heap.Pop(&this.queue).([]int)
		if has(this.included, top[0], top[1]) {
			prices = append(prices, this.remove(top[0], top[1]))
			out = append(out, top[:2])
		}
	}
	for i, it := range out {
		this.add([]int{it[0], it[1], prices[i]})
	}
	return out
}

type unrented struct {
	included map[int]map[int]int
	queue    map[int]hp.Heap[[]int]
}

func newUnrented(n int, entries [][]int) *unrented {
	included := make(map[int]map[int]int, n)
	queue := map[int]hp.Heap[[]int]{}
	for _, it := range entries {
		shop, movie, price := it[0], it[1], it[2]
		if _, ok := included[shop]; ok {
			included[shop][movie] = price
		} else {
			included[shop] = map[int]int{
				movie: price,
			}
		}
		if h, ok := queue[movie]; ok {
			heap.Push(&h, []int{shop, price})
			queue[movie] = h
		} else {
			queue[movie] = hp.New[[]int]([][]int{{shop, price}}, func(a, b []int) bool {
				return a[1] < b[1] || (a[1] == b[1] && a[0] < b[0])
			})
		}
	}

	out := &unrented{
		included: included,
		queue:    queue,
	}
	return out
}

func (this *unrented) add(item []int) {
	add(this.included, item)
	if h, ok := this.queue[item[1]]; ok {
		heap.Push(&h, []int{item[0], item[2]})
		this.queue[item[1]] = h
	}
}

func (this *unrented) remove(shop, movie int) int {
	price := this.included[shop][movie]
	delete(this.included[shop], movie)
	return price
}

func (this *unrented) getTop5(movie int) []int {
	if h, ok := this.queue[movie]; ok {
		out := make([]int, 0, 5)
		prices := make([]int, 0, 5)
		for len(out) < 5 && h.Len() > 0 {
			top := heap.Pop(&h).([]int)
			if has(this.included, top[0], movie) {
				prices = append(prices, this.remove(top[0], movie))
				out = append(out, top[0])
			}
		}
		this.queue[movie] = h
		for i, it := range out {
			this.add([]int{it, movie, prices[i]})
		}
		return out
	}
	return nil
}

type MovieRentingSystem struct {
	unrented *unrented
	rented   *rented
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	return MovieRentingSystem{
		unrented: newUnrented(n, entries),
		rented:   newRented(n),
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	return this.unrented.getTop5(movie)
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	price := this.unrented.remove(shop, movie)
	this.rented.add([]int{shop, movie, price})
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	price := this.rented.remove(shop, movie)
	this.unrented.add([]int{shop, movie, price})
}

func (this *MovieRentingSystem) Report() [][]int {
	return this.rented.getTop5()
}
