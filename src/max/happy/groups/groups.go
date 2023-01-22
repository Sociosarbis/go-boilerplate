package groups

func get(freq int64, i int) int64 {
	return freq & (31 << (i * 5)) >> (i * 5)
}

func set(freq int64, i int, v int64, mask int64) int64 {
	return freq&(mask-(31<<(i*5))) | (v << (i * 5))
}

func maxHappyGroups(batchSize int, groups []int) int {
	if batchSize == 1 {
		return len(groups)
	}
	var freq int64 = 0
	ret := 1
	var mask int64 = (1 << (5 * batchSize)) - 1
	for _, group := range groups {
		if group%batchSize == 0 {
			ret++
		} else {
			i := group % batchSize
			freq = set(freq, i, get(freq, i)+1, mask)
		}
	}
	dp := map[int64][]int{
		freq: make([]int, batchSize),
	}
	dp[freq][0] = ret
	for {
		nextDp := make(map[int64][]int, 0)
		for k, v := range dp {
			for i := 1; i < batchSize; i++ {
				c := get(k, i)
				if c != 0 {
					nextK := set(k, i, c-1, mask)
					if _, ok := nextDp[nextK]; !ok {
						nextDp[nextK] = make([]int, batchSize)
					}
					for j := 0; j < batchSize; j++ {
						if v[j] != 0 {
							index := (i + j) % batchSize
							var nextV int
							if index != 0 {
								nextV = v[j]
							} else {
								nextV = v[j] + 1
							}
							if nextV > nextDp[nextK][index] {
								nextDp[nextK][index] = nextV
							}
						}
					}
				}
			}
		}
		dp = nextDp
		if _, ok := nextDp[0]; ok {
			break
		}
	}
	ret = 0
	for i, c := range dp[0] {
		if i == 0 {
			if c-1 > ret {
				ret = c - 1
			}
		} else {
			if c > ret {
				ret = c
			}
		}
	}
	return ret
}
