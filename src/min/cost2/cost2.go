package cost2

func minCost(nums []int, x int) int64 {
	n := len(nums)
	var ret int64
	h := make([]int, n)
	copy(h, nums)

	for i := 1; i <= n; i++ {
		var cost int64 = int64(x) * int64(i-1)
		for j := range nums {
			if nums[(j+i-1)%n] < h[j] {
				h[j] = nums[(j+i-1)%n]
			}
			cost += int64(h[j])
		}
		if i == 1 {
			ret = cost
		} else if cost < ret {
			ret = cost
		}
	}
	return ret
}
