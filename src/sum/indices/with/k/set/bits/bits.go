package bits

func sumIndicesWithKSetBits(nums []int, k int) int {
	ret := 0
	for i, num := range nums {
		if getBits(i) == k {
			ret += num
		}
	}
	return ret
}

func getBits(num int) int {
	ret := 0

	for num != 0 {
		ret++
		num &= num - 1
	}
	return ret
}
