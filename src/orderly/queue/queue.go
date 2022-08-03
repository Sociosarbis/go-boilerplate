package queue

import "sort"

type bytes []byte

func (this bytes) Len() int {
	return len(this)
}

func (this bytes) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this bytes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func orderlyQueue(s string, k int) string {
	// 只要k大于2就能做到两两比较，把较小的那个放到后尾，循环多次总会有序
	if k > 1 {
		old := []byte(s)
		sort.Sort(bytes(old))
		return string(old)
	} else {
		ret := 0
		for i := 1; i < len(s); i++ {
			for j := 0; j < len(s); j++ {
				i1 := (ret + j) % len(s)
				i2 := (i + j) % len(s)
				if s[i1] < s[i2] {
					break
				} else if s[i1] > s[i2] {
					ret = i
					break
				}
			}
		}
		return s[ret:] + s[:ret]
	}
}
