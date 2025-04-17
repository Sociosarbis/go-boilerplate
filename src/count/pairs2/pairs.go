package pairs2

func countPairs(nums []int, k int) int {
	n := len(nums)
	var ret int
	for i := 0; i < n; i++ {
		a := nums[i]
		for j := i + 1; j < n; j++ {
			if a == nums[j] && (i*j)%k == 0 {
				ret++
			}
		}
	}
	return ret
}
