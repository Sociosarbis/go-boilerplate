package sum2

import "github.com/boilerplate/src/tree/avl"

type item struct {
	counter map[int]int
	num     int
}

func (self item) Lt(other item) bool {
	if c1, ok := self.counter[self.num]; ok {
		if c2, ok := other.counter[other.num]; ok {
			if c1 < c2 || (c1 == c2 && self.num < other.num) {
				return true
			}
		}
	}
	return false
}

func (self item) Gt(other item) bool {
	if c1, ok := self.counter[self.num]; ok {
		if c2, ok := other.counter[other.num]; ok {
			if c1 > c2 || (c1 == c2 && self.num > other.num) {
				return true
			}
		}
	}
	return false
}

func (self item) Eq(other item) bool {
	return self.num == other.num
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums) - k + 1
	out := make([]int64, n)
	counter := make(map[int]int, len(nums))
	var left *avl.AvlNode[item]
	var right *avl.AvlNode[item]
	var l1 int
	for i := 0; i < k; i++ {
		if c, ok := counter[nums[i]]; ok {
			counter[nums[i]] = c + 1
		} else {
			counter[nums[i]] = 1
		}
	}
	for k, v := range counter {
		left = left.Insert(item{
			counter: counter,
			num:     k,
		})
		out[0] += int64(k) * int64(v)
	}
	l1 = len(counter)
	for l1 > x {
		node := left.GetMinValueNode().GetValue()
		left = left.Delete(node)
		right = right.Insert(node)
		out[0] -= int64(node.num) * int64(counter[node.num])
		l1--
	}
	for i := 1; i < n; i++ {
		n1 := item{
			counter: counter,
			num:     nums[i-1],
		}
		out[i] = out[i-1]
		a := left.Search(n1)
		if a == nil {
			a = right.Search(n1)
			if a != nil {
				right = right.Delete(n1)
				if c, ok := counter[n1.num]; ok {
					counter[n1.num] = c - 1
				}
				right = right.Insert(n1)
			}
		} else {
			left = left.Delete(n1)
			if c, ok := counter[n1.num]; ok {
				counter[n1.num] = c - 1
			}
			out[i] -= int64(n1.num)
			left = left.Insert(n1)
		}
		n2 := item{
			counter: counter,
			num:     nums[i+k-1],
		}
		a = left.Search(n2)
		if a == nil {
			a = right.Search(n2)
			if a != nil {
				right = right.Delete(n2)
			}
			if c, ok := counter[n2.num]; ok {
				counter[n2.num] = c + 1
			} else {
				counter[n2.num] = 1
			}
			right = right.Insert(n2)
		} else {
			left = left.Delete(n2)
			if c, ok := counter[n2.num]; ok {
				counter[n2.num] = c + 1
			}
			out[i] += int64(n2.num)
			left = left.Insert(n2)
		}
		for l1 < x && right != nil {
			node := right.GetMaxValueNode().GetValue()
			right = right.Delete(node)
			left = left.Insert(node)
			out[i] += int64(node.num) * int64(counter[node.num])
			l1++
		}
		for right != nil {
			r := right.GetMaxValueNode().GetValue()
			l := left.GetMinValueNode().GetValue()
			if r.Gt(l) {
				out[i] += int64(r.num)*int64(counter[r.num]) - int64(l.num)*int64(counter[l.num])
				left = left.Delete(l).Insert(r)
				right = right.Delete(r).Insert(l)
			} else {
				break
			}
		}
	}
	return out
}
