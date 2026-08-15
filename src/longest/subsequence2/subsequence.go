package subsequence2

func longestSubsequence(nums []int) int {
	var hasNonZero bool
	var temp int
	for _, num := range nums {
		if !hasNonZero && num != 0 {
			hasNonZero = true
		}
		temp ^= num
	}
	if temp != 0 {
		return len(nums)
	}
	if hasNonZero {
		return len(nums) - 1
	}
	return 0
}
