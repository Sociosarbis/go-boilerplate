package units

import "sort"

type BoxTypes [][]int

func (this BoxTypes) Len() int {
	return len(this)
}

func (this BoxTypes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this BoxTypes) Less(i, j int) bool {
	return this[i][1] > this[j][1]
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Sort(BoxTypes(boxTypes))
	i := 0
	ret := 0
	for i < len(boxTypes) {
		if boxTypes[i][0] >= truckSize {
			ret += truckSize * boxTypes[i][1]
		} else {
			truckSize -= boxTypes[i][0]
			ret += boxTypes[i][0] * boxTypes[i][1]
		}
		i++
	}
	return ret
}
