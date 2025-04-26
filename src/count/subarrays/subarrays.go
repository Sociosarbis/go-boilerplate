package subarrays

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func count(n int) int64 {
	return int64(n+1) * int64(n) / 2
}

func consume(pivots [][2]int, l, r int) int64 {
	ret := count(r - l)
	n := len(pivots)
	for i, p := range pivots {
		if p[0] == 3 {
			if i == 0 {
				ret -= count(p[1] - l)
			} else {
				ret += count(p[1] - l)
			}
			end := r
			if i+1 < n {
				end = pivots[i+1][1]
			}
			ret -= count(end - p[1] - 1)
		} else {
			end := r
			if i > 0 && pivots[i-1][0] != p[0] {
				ret += count(p[1] - l)
			}
			if i+1 < n {
				if pivots[i+1][0] == p[0] {
					continue
				}
				end = pivots[i+1][1]
			}
			ret -= count(end - l)
		}
		l = p[1] + 1
	}
	return ret
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var i, j int
	pivots := [][2]int{}
	n := len(nums)
	var ret int64
	var s int
	for j < n {
		if nums[j] < minK || nums[j] > maxK {
			if s == 3 {
				ret += consume(pivots, i, j)
			}
			s = 0
			pivots = pivots[:0]
			j++
			i = j
		} else {
			var t int
			if nums[j] == minK {
				t |= 1
			}
			if nums[j] == maxK {
				t |= 2
			}
			s |= t
			pivots = append(pivots, [2]int{t, j})
			j++
		}
	}
	if s == 3 {
		ret += consume(pivots, i, j)
	}
	return ret
}
