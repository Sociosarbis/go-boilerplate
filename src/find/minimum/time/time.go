package time

import (
	"sort"
)

func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	n := len(tasks)
	take_times := [][2]int{}
loop:
	for i := 0; i < n; i++ {
		start, end, dur := tasks[i][0], tasks[i][1], tasks[i][2]
		for _, time := range take_times {
			if start <= time[1] {
				cur_dur := time[1] - time[0] + 1
				overlapped := time[1] - start + 1
				if overlapped > cur_dur {
					overlapped = cur_dur
				}
				if overlapped >= dur {
					continue loop
				} else {
					dur -= overlapped
				}
			}
		}
		start = end - dur + 1
		for j := len(take_times) - 1; j >= 0; j-- {
			time := take_times[j]
			if time[1] >= start {
				start = time[0] - (time[1] - start + 1)
				take_times = take_times[:j]
			} else {
				break
			}
		}
		take_times = append(take_times, [2]int{start, end})
	}
	var ret int
	for _, time := range take_times {
		ret += time[1] - time[0] + 1
	}
	return ret
}
