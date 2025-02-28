package ratings

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type FoodRatings struct {
	buckets      map[string]*h.Heap[pair]
	foodToCusine map[string]string
	foodToRating map[string]int
}

type pair struct {
	a string
	b int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	n := len(cuisines)
	m := map[string][]pair{}
	foodToCusine := make(map[string]string, n)
	foodToRating := make(map[string]int, n)
	for i := 0; i < n; i++ {
		if _, ok := m[cuisines[i]]; ok {
			m[cuisines[i]] = append(m[cuisines[i]], pair{foods[i], ratings[i]})
		} else {
			m[cuisines[i]] = []pair{{foods[i], ratings[i]}}
		}
		foodToCusine[foods[i]] = cuisines[i]
		foodToRating[foods[i]] = ratings[i]
	}
	hpm := make(map[string]*h.Heap[pair], len(m))
	for cuisine, pairs := range m {
		bucket := h.New[pair](pairs, func(a, b pair) bool {
			if a.b > b.b {
				return true
			} else if a.b < b.b {
				return false
			}
			l1 := len(a.a)
			l2 := len(b.a)
			var l int
			if l1 < l2 {
				l = l1
			} else {
				l = l2
			}
			for i := 0; i < l; i++ {
				if a.a[i] < b.a[i] {
					return true
				} else if a.a[i] > b.a[i] {
					return false
				}
			}
			return l1 < l2
		})
		hpm[cuisine] = &bucket
	}
	return FoodRatings{buckets: hpm, foodToCusine: foodToCusine, foodToRating: foodToRating}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	if cuisine, ok := this.foodToCusine[food]; ok {
		heap.Push(this.buckets[cuisine], pair{a: food, b: newRating})
	}
	this.foodToRating[food] = newRating
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	for {
		if bucket, ok := this.buckets[cuisine]; ok {
			top := bucket.Top()
			if rating, ok := this.foodToRating[top.a]; ok && rating == top.b {
				return top.a
			}
			heap.Pop(bucket)
		} else {
			return ""
		}
	}
}
