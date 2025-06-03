package candies

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	bfs := make([]int, 0, n)
	boxes := make([]bool, n)
	var ret int
	for _, i := range initialBoxes {
		boxes[i] = true
		if status[i] == 1 {
			bfs = append(bfs, i)
		}
	}
	n = len(bfs)
	var i int
	for i < n {
		for ; i < n; i++ {
			ret += candies[bfs[i]]
			for _, box := range containedBoxes[bfs[i]] {
				if !boxes[box] {
					boxes[box] = true
					if status[box] == 1 {
						bfs = append(bfs, box)
					}
				}
			}
			for _, key := range keys[bfs[i]] {
				if status[key] != 1 {
					status[key] = 1
					if boxes[key] {
						bfs = append(bfs, key)
					}
				}
			}
		}
		i, n = n, len(bfs)
	}
	return ret
}
