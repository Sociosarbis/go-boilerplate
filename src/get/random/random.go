package random

import (
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	values []int
}

func Constructor(head *ListNode) Solution {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return Solution{
		values,
	}
}

func (this *Solution) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}
