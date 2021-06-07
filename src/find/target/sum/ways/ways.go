package ways

func findTargetSumWays(nums []int, target int) int {
	dp := map[int]int{
		0: 1,
	}
	for _, num := range nums {
		newDp := map[int]int{}
		for key, val := range dp {
			n1 := key + num
			n2 := key - num
			count1, exist1 := newDp[n1]
			if !exist1 {
				newDp[n1] = val
			} else {
				newDp[n1] = count1 + val
			}

			count2, exist2 := newDp[n2]

			if !exist2 {
				newDp[n2] = val
			} else {
				newDp[n2] = count2 + val
			}
		}
		dp = newDp
	}
	ret := 0
	for key, count := range dp {
		if key == target {
			ret += count
		}
	}
	return ret
}
