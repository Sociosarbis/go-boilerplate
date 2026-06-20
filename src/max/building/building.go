package building

import "sort"

func getMax(distance, from, to int) (int, int) {
	temp := from + distance
	if temp <= to {
		return temp, temp
	}
	a := (distance + to - from) / 2
	return from + a, to
}

func maxBuilding(n int, restrictions [][]int) int {
	if len(restrictions) == 0 {
		return n - 1
	}
	if len(restrictions) > 1 {
		sort.Slice(restrictions, func(i, j int) bool {
			return restrictions[i][0] < restrictions[j][0]
		})
	}
	m := len(restrictions)
	if restrictions[m-1][0] != n {
		restrictions = append(restrictions, []int{n, n - 1})
		m++
	}
	for i := m - 1; i > 0; i-- {
		if restrictions[i-1][1]-restrictions[i][1] > restrictions[i][0]-restrictions[i-1][0] {
			restrictions[i-1][1] = restrictions[i][1] + restrictions[i][0] - restrictions[i-1][0]
		}
	}
	max, current := getMax(restrictions[0][0]-1, 0, restrictions[0][1])
	out := max
	for i := 1; i < m; i++ {
		max, current = getMax(restrictions[i][0]-restrictions[i-1][0], current, restrictions[i][1])
		if max > out {
			out = max
		}
	}
	return out
}
