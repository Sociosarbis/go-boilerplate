package assign

import (
	"sort"

	"github.com/boilerplate/src/tree/avl"
)

func check(tasks []int, workers []int, pills, strength int) bool {
	n := len(tasks)
	m := len(workers)
	var tree *avl.AvlNode[avl.Integer]
	for i := m - 1; i >= m-n; i-- {
		tree = tree.Insert(avl.Integer(workers[i]))
	}

	for i := 0; i < n; i++ {
		value := tasks[n-i-1]
		res := tree.Search(avl.Integer(value))
		if res != nil {
			tree = tree.Delete(res.GetValue())
		} else {
			if pills > 0 {
				value -= strength
				pills--
			} else {
				return false
			}
			res := tree.Search(avl.Integer(value))
			if res == nil {
				return false
			}
			tree = tree.Delete(res.GetValue())
		}
	}
	return true
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	n := len(tasks)
	m := len(workers)
	var l int
	r := n
	if r > m {
		r = m
	}
	var ret int
	for l <= r {
		mid := (l + r) >> 1
		if check(tasks[:mid], workers, pills, strength) {
			if mid > ret {
				ret = mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
