package score

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func maxScore(nums []int) int {
	n := len(nums)
	round := n / 2
	dp := map[int]int{
		0: 0,
	}
	for i := 0; i < round; i++ {
		nextDp := map[int]int{}
		for prevMask, v := range dp {
			for j := 0; j < n; j++ {
				if prevMask&(1<<j) == 0 {
					for k := j + 1; k < n; k++ {
						if prevMask&(1<<k) == 0 {
							nextV := v + gcd(nums[j], nums[k])*(round-i)
							nextMask := prevMask | (1 << j) | (1 << k)
							if prevV, ok := nextDp[nextMask]; ok {
								if prevV < nextV {
									nextDp[nextMask] = nextV
								}
							} else {
								nextDp[nextMask] = nextV
							}
						}
					}
				}
			}
		}
		dp = nextDp
	}
	ret := 0
	for _, v := range dp {
		if v > ret {
			ret = v
		}
	}
	return ret
}
