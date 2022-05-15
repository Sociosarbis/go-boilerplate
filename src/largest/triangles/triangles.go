package triangles

import "math"

func dfs(indices *[]int, points *[][]int) float64 {
	n := len(*indices)
	if n == 3 {
		a := (*indices)[0]
		b := (*indices)[1]
		c := (*indices)[2]
		x1 := (*points)[b][0] - (*points)[a][0]
		y1 := (*points)[b][1] - (*points)[a][1]
		x2 := (*points)[c][0] - (*points)[a][0]
		y2 := (*points)[c][1] - (*points)[a][1]
		return math.Abs(float64(x1*y2-y1*x2)) * 0.5
	}
	i := (*indices)[n-1]
	r := len(*points) + n - 2
	var ret float64
	for j := i + 1; j < r; j++ {
		*indices = append(*indices, j)
		temp := dfs(indices, points)
		*indices = (*indices)[:n]
		if temp > ret {
			ret = temp
		}
	}
	return ret
}

func largestTriangleArea(points [][]int) float64 {
	var ret float64
	indices := []int{}
	for i := 0; i < len(points)-2; i++ {
		indices = append(indices, i)
		temp := dfs(&indices, &points)
		if temp > ret {
			ret = temp
		}
		indices = indices[:0]
	}
	return ret
}
