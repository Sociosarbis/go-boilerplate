package difference

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	ret := [][]int{{arr[0], arr[1]}}
	min := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
			ret = [][]int{{arr[i-1], arr[i]}}
		} else if arr[i]-arr[i-1] == min {
			ret = append(ret, []int{arr[i-1], arr[i]})
		}
	}
	return ret
}
