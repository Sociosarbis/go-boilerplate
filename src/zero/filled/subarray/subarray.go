package subarray

func zeroFilledSubarray(nums []int) int64 {
	var out int64
	prev := -1
	for i, num := range nums {
		if num == 0 {
			if prev == -1 {
				prev = i
			}
		} else {
			if prev != -1 {
				out += int64(1+i-prev) * int64(i-prev) / 2
				prev = -1
			}
		}
	}
	if prev != -1 {
		return out + int64(1+len(nums)-prev)*int64(len(nums)-prev)/2
	}
	return out
}
