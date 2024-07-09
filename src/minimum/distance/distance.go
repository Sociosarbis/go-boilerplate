package distance

import "sort"

func minimumDistance(points [][]int) int {
	n := len(points)
	transformed := make([][2]int, n)

	for i, point := range points {
		// 距离计算所需的元素，都可以用这两个数来表示
		transformed[i][0], transformed[i][1] = point[0]+point[1], point[0]-point[1]
	}

	sortByX := make([]int, n)
	sortByY := make([]int, n)

	for i := 0; i < n; i++ {
		sortByX[i] = i
		sortByY[i] = i
	}
	sort.Slice(sortByX, func(i, j int) bool {
		return transformed[sortByX[i]][0] < transformed[sortByX[j]][0]
	})
	sort.Slice(sortByY, func(i, j int) bool {
		return transformed[sortByY[i]][1] < transformed[sortByY[j]][1]
	})

	var ret int = 4 * 1e8
	for _, option := range [4]int{sortByX[0], sortByX[n-1], sortByY[0], sortByY[n-1]} {
		var minX int
		if option == sortByX[0] {
			minX = transformed[sortByX[1]][0]
		} else {
			minX = transformed[sortByX[0]][0]
		}
		var maxX int
		if option == sortByX[n-1] {
			maxX = transformed[sortByX[n-2]][0]
		} else {
			maxX = transformed[sortByX[n-1]][0]
		}

		var minY int
		if option == sortByY[0] {
			minY = transformed[sortByY[1]][1]
		} else {
			minY = transformed[sortByY[0]][1]
		}
		var maxY int
		if option == sortByY[n-1] {
			maxY = transformed[sortByY[n-2]][1]
		} else {
			maxY = transformed[sortByY[n-1]][1]
		}
		var temp int
		if maxX-minX > maxY-minY {
			temp = maxX - minX
		} else {
			temp = maxY - minY
		}
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
