package string

import "github.com/boilerplate/src/heap"

func repeatLimitedString(s string, repeatLimit int) string {
	dict := make([]int, 26)

	for _, c := range s {
		dict[c-97] += 1
	}

	h := heap.New([][2]int{}, func(a [2]int, b [2]int) bool {
		return a[0] < b[0]
	})

	for i, count := range dict {
		if count > 0 {
			h.Push([2]int{97 + i, count})
		}
	}

	ret := make([]byte, 0, len(s))
	last := 0
	count := 0
	for h.Len() != 0 {
		top1 := h.Pop().([2]int)
		if top1[0] == last && count == repeatLimit {
			if h.Len() != 0 {
				top2 := h.Pop().([2]int)
				ret = append(ret, byte(top2[0]))
				top2[1]--
				count = 1
				last = top2[0]
				if top2[1] != 0 {
					h.Push(top2)
				}
			} else {
				break
			}
			h.Push(top1)
		} else {
			if top1[0] != last {
				count = 1
			} else {
				count++
			}
			ret = append(ret, byte(top1[0]))
			top1[1]--
			last = top1[0]
			if top1[1] != 0 {
				h.Push(top1)
			}
		}
	}

	return string(ret)
}
