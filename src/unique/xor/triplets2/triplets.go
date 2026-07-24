package triplets2

func uniqueXorTriplets(nums []int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	count := 1
	for count <= max {
		count <<= 1
	}
	twoXors := make([]bool, count)
	threeXors := make([]bool, count)
	n := len(nums)
	var out int
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		if !threeXors[num] {
			threeXors[num] = true
			out++
		}
		for j := i + 1; j < n; j++ {
			xor := num ^ nums[j]
			twoXors[xor] = true
		}
		for j := 0; j < count; j++ {
			if !twoXors[j] {
				continue
			}
			xor := num ^ j
			if !threeXors[xor] {
				threeXors[xor] = true
				out++
			}
		}
	}
	return out
}
