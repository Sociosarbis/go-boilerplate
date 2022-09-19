package sort

import "sort"

type Nums struct {
	counter []int
	data    []int
}

func (this Nums) Len() int {
	return len(this.data)
}

func (this Nums) Swap(i, j int) {
	this.data[i], this.data[j] = this.data[j], this.data[i]
}

func (this Nums) Less(i, j int) bool {
	a := this.data[i]
	b := this.data[j]
	if this.counter[a+100] < this.counter[b+100] || (this.counter[a+100] == this.counter[b+100] && a > b) {
		return true
	}
	return false
}

func frequencySort(nums []int) []int {
	counter := make([]int, 201)
	for _, num := range nums {
		counter[num+100]++
	}
	subject := Nums{
		counter: counter,
		data:    nums,
	}
	sort.Sort(subject)
	return subject.data
}
