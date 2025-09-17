package containers

import (
	"sort"
)

type NumberContainers struct {
	indexToNumber   map[int]int
	numberToIndices map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexToNumber:   make(map[int]int, 1e5),
		numberToIndices: make(map[int][]int, 1e5),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if num, ok := this.indexToNumber[index]; ok {
		if num == number {
			return
		}
		if indices, ok := this.numberToIndices[num]; ok {
			i, found := sort.Find(len(indices), func(i int) int {
				return index - indices[i]
			})
			if found {
				copy(indices[i:], indices[i+1:])
				indices = indices[:len(indices)-1]
				if len(indices) == 0 {
					delete(this.numberToIndices, num)
				} else {
					this.numberToIndices[num] = indices
				}
			}
		}
	}
	this.indexToNumber[index] = number
	if indices, ok := this.numberToIndices[number]; ok {
		i := sort.Search(len(indices), func(i int) bool {
			return indices[i] >= index
		})
		if i >= len(indices) {
			this.numberToIndices[number] = append(indices, index)
		} else {
			indices = append(indices, 0)
			copy(indices[i+1:], indices[i:])
			indices[i] = index
			this.numberToIndices[number] = indices
		}
	} else {
		this.numberToIndices[number] = []int{index}
	}

}

func (this *NumberContainers) Find(number int) int {
	if indices, ok := this.numberToIndices[number]; ok {
		return indices[0]
	}
	return -1
}
