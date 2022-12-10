package height

import (
	"sort"
)

type Cuboids [][]int

func (this Cuboids) Len() int {
	return len(this)
}

func (this Cuboids) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Cuboids) Less(i, j int) bool {
	return this[i][2] < this[j][2] || (this[i][2] == this[j][2] && (this[i][1] < this[j][1] || (this[i][1] == this[j][1] && this[i][0] < this[j][0])))
}

func maxHeight(cuboids [][]int) int {
	for _, cuboid := range cuboids {
		sort.Ints(cuboid)
	}
	// 因为是求最高的组合，而且能叠的要求是长宽高都不大于上一个，并可以任意旋转，所以以最大的边为高是等效的，
	// 并需要以高为排序的第一尺度
	sort.Sort(Cuboids(cuboids))
	dp1 := make([]int, len(cuboids))
	dp2 := make([]int, len(cuboids))
	for i, cuboid := range cuboids {
		dp1[i] = cuboid[2]
		for j := i - 1; j >= 0; j-- {
			if cuboid[1] >= cuboids[j][1] && cuboid[0] >= cuboids[j][0] {
				if cuboid[2]+dp1[j] > dp1[i] {
					dp1[i] = cuboid[2] + dp1[j]
				}
				if dp1[j] == dp2[j] {
					break
				}
			}
		}

		if i == 0 || dp1[i] > dp2[i-1] {
			dp2[i] = dp1[i]
		} else {
			dp2[i] = dp2[i-1]
		}
	}
	return dp2[len(cuboids)-1]
}
