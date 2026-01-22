package removal

import (
	"container/heap"

	h "github.com/boilerplate/src/heap"
)

type node struct {
	prev  *node
	next  *node
	index int
	value int
}

type item struct {
	first  *node
	second *node
	total  int
}

func minimumPairRemoval(nums []int) int {
	head := &node{
		index: 0,
		value: nums[0],
	}
	cur := head
	n := len(nums)
	hp := h.New([]item{}, func(a, b item) bool {
		return a.total < b.total || (a.total == b.total && a.first.index < b.first.index)
	})
	var count int
	for i := 1; i < n; i++ {
		next := &node{
			index: i,
			value: nums[i],
			prev:  cur,
		}
		if cur.value > next.value {
			count++
		}
		cur.next = next
		heap.Push(&hp, item{
			first:  cur,
			second: next,
			total:  cur.value + next.value,
		})
		cur = next
	}
	cur = head
	var out int
	for count > 0 {
		top := heap.Pop(&hp).(item)
		if top.first.next == top.second {
			out++
			mergedNode := &node{
				index: top.first.index,
				value: top.total,
				prev:  top.first.prev,
				next:  top.second.next,
			}
			if top.first.value > top.second.value {
				count--
			}
			if top.first.prev != nil {
				if top.first.prev.value > top.first.value {
					count--
				}
				if top.first.prev.value > mergedNode.value {
					count++
				}
				top.first.prev.next = mergedNode
				heap.Push(&hp, item{
					total:  top.first.prev.value + mergedNode.value,
					first:  top.first.prev,
					second: mergedNode,
				})
			}
			if top.second.next != nil {
				if top.second.value > top.second.next.value {
					count--
				}
				if mergedNode.value > top.second.next.value {
					count++
				}
				top.second.next.prev = mergedNode
				heap.Push(&hp, item{
					total:  top.second.next.value + mergedNode.value,
					first:  mergedNode,
					second: top.second.next,
				})
			}
			top.first.prev, top.first.next, top.second.prev, top.second.next = nil, nil, nil, nil
		}
	}
	return out
}
