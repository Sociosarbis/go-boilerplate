package moves2

func minimumMoves(nums []int, k int, maxChanges int) int64 {
	n := len(nums)
	var c int
	pos := []int{}
	for i, num := range nums {
		if num == 1 {
			pos = append(pos, i)
			if c < 1 {
				c = 1
			}
			if i+1 < n && nums[i+1] == 1 {
				if c < 2 {
					c = 2
				}
				if i+2 < n && nums[i+2] == 1 {
					if c < 3 {
						c = 3
					}
				}
			}
		}
	}
	if k <= c {
		return int64(k - 1)
	}
	if c+maxChanges >= k {
		remain := k - c
		if c <= 1 {
			return int64(2 * remain)
		} else {
			return int64(2*remain + c - 1)
		}
	}
	var ret int64 = 1e10
	k -= maxChanges
	m := len(pos)
	prefixSum := make([]int64, m+1)
	for i, p := range pos {
		prefixSum[i+1] = prefixSum[i] + int64(p)
	}
	var l int
	var r int
	/** 尝试各种长度为k的子数组 */
	for ; r < m; r++ {
		if r-l+1 == k {
			mid := (l + r) / 2
			temp := int64(mid-l)*int64(pos[mid]) - (prefixSum[mid] - prefixSum[l]) + (prefixSum[r+1] - prefixSum[mid+1]) - int64(r-mid)*int64(pos[mid])
			if temp < ret {
				ret = temp
			}
			l++
		}
	}
	return ret + int64(maxChanges*2)
}
