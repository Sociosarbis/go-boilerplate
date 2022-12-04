package cost

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ret := baseCosts[0]
	var temp int
	if ret > target {
		temp = ret - target
	} else {
		temp = target - ret
	}
	for _, b := range baseCosts {
		if b == target {
			return target
		} else if b < target {
			if target-b < temp || (target-b == temp && b < ret) {
				temp = target - b
				ret = b
			}
		} else {
			if b-target < temp {
				temp = b - target
				ret = b
			}
		}
		dp := map[int]struct{}{
			b: {},
		}
		for _, t := range toppingCosts {
			oldKeys := make([]int, len(dp))
			i := 0
			for k := range dp {
				oldKeys[i] = k
				i++
			}
			for _, k := range oldKeys {
				for i := 1; i <= 2; i++ {
					value := t*i + k
					if _, ok := dp[value]; !ok {
						if value == target {
							return target
						} else if value < target {
							if target-value < temp || (target-value == temp && value < ret) {
								temp = target - value
								ret = value
							}
							dp[value] = struct{}{}
						} else {
							if value-target < temp {
								temp = value - target
								ret = value
							}
						}
					}
				}
			}
		}
	}
	return ret
}
