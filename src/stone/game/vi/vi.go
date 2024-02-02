package vi

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	values := make([][2]int, n)
	for i := 0; i < n; i++ {
		values[i] = [2]int{aliceValues[i] + bobValues[i], i}
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i][0] < values[j][0]
	})
	temp := 0
	for i := n - 1; i >= 0; i -= 2 {
		temp += aliceValues[values[i][1]]
	}
	for i := n - 2; i >= 0; i -= 2 {
		temp -= bobValues[values[i][1]]
	}
	if temp > 0 {
		return 1
	} else if temp < 0 {
		return -1
	}
	return 0
}
