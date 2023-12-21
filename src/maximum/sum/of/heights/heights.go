package heights

func maximumSumOfHeights(maxHeights []int) int64 {
	dp := make([]int64, len(maxHeights))
	iters := [2][3]int{{len(maxHeights) - 1, -1, -1}, {0, len(maxHeights), 1}}
	var ret int64
	for ii, iter := range iters {
		heights := make([]int, 0, len(maxHeights))
		counter := map[int]int{}
		var prevVal int64
		for i := iter[0]; i != iter[1]; i += iter[2] {
			height := maxHeights[i]
			prevVal += int64(height)
			oldCount, ok := counter[height]
			if !ok {
				counter[height] = 1
			} else {
				counter[height] = oldCount + 1
			}
			j := len(heights)
			for j > 0 && heights[j-1] > height {
				j--
				diffCount := counter[heights[j]]
				counter[height] += diffCount
				prevVal += int64(diffCount) * int64(height-heights[j])
				delete(counter, heights[j])
			}
			heights = heights[:j]
			if j == 0 || heights[j-1] != height {
				heights = append(heights, height)
			}
			if ii == 1 {
				nextValue := dp[i] + prevVal - int64(height)
				if nextValue > ret {
					ret = nextValue
				}
			} else {
				dp[i] = prevVal
			}
		}
	}
	return ret
}
