package pick

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	ret := Solution{
		nums,
	}

	return ret
}

func (this *Solution) Pick(target int) int {
	count := 0
	ret := 0
	for i, num := range this.nums {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				ret = i
			}
		}
	}
	return ret
}
