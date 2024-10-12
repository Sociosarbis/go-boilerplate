package xor

func duplicateNumbersXOR(nums []int) int {
	var mask int64
	var ret int
	for _, num := range nums {
		if mask&(1<<num) != 0 {
			ret ^= num
		} else {
			mask |= 1 << num
		}
	}
	return ret
}
