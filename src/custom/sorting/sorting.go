package sorting

import "sort"

type item struct {
	c byte
	o int
}

type items []item

func (this items) Len() int {
	return len(this)
}

func (this items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this items) Less(i, j int) bool {
	return this[i].o > this[j].o
}

func customSortString(order string, s string) string {
	orders := make([]int, 26)
	for i, c := range order {
		orders[c-97] = len(s) - i
	}
	arr := make(items, len(s))
	for i, c := range s {
		arr[i] = item{
			byte(c),
			orders[c-97],
		}
	}
	sort.Sort(arr)
	ret := make([]byte, len(s))
	for i, item := range arr {
		ret[i] = item.c
	}
	return string(ret)
}
