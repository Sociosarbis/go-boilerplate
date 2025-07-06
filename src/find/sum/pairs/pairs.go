package pairs

type FindSumPairs struct {
	nums1   []int
	nums2   []int
	counter map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	counter := make(map[int]int, len(nums2))
	for _, num := range nums2 {
		if c, ok := counter[num]; ok {
			counter[num] = c + 1
		} else {
			counter[num] = 1
		}
	}
	return FindSumPairs{
		nums1,
		nums2,
		counter,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	prev := this.nums2[index]
	this.nums2[index] += val
	if num, ok := this.counter[prev]; ok {
		this.counter[prev] = num - 1
	}
	if num, ok := this.counter[this.nums2[index]]; ok {
		this.counter[this.nums2[index]] = num + 1
	} else {
		this.counter[this.nums2[index]] = 1
	}
}

func (this *FindSumPairs) Count(tot int) int {
	var out int
	for _, num := range this.nums1 {
		if c, ok := this.counter[tot-num]; ok {
			out += c
		}
	}
	return out
}
