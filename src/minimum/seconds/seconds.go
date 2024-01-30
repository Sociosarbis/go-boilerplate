package seconds

import "github.com/boilerplate/src/utils"

func minimumSeconds(nums []int) int {
	visited := map[int][]int{}
	n := len(nums)
	ret := n
	for i, num := range nums {
		if values, ok := visited[num]; !ok {
			visited[num] = []int{i, i, 0}
		} else {
			temp := (i - values[1]) / 2
			if temp > values[2] {
				values[2] = temp
			}
			values[1] = i
			visited[num] = values
		}
	}
	for _, values := range visited {
		temp := utils.Max((values[0]+n-values[1])/2, values[2])
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
