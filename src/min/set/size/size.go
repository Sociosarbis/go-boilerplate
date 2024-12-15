package size

import "sort"

func minSetSize(arr []int) int {
	n := len(arr)
	counter := make(map[int]int, n)
	for _, num := range arr {
		if c, ok := counter[num]; ok {
			counter[num] = c + 1
		} else {
			counter[num] = 1
		}
	}
	queue := make([][2]int, 0, len(counter))
	for k, v := range counter {
		queue = append(queue, [2]int{k, v})
	}
	sort.Slice(queue, func(i, j int) bool {
		return queue[i][1] > queue[j][1]
	})
	n = n / 2
	var ret int
	for n > 0 {
		n -= queue[ret][1]
		ret++
	}
	return ret
}
