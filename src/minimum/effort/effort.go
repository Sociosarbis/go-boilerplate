package effort

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
	})
	var out, temp int
	for _, task := range tasks {
		value := max(task[0], task[1])
		if temp < value {
			out += value - temp
			temp += value - temp
		}
		temp -= task[0]
	}
	return out
}
