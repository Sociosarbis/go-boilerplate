package sum

func subsetXORSum(nums []int) int {
	cs := [2][32]int{}
	cs[0][0] = 1
	var ret int
	for i, num := range nums {
		index := i & 1
		copy(cs[1-index][:], cs[index][:])
		for b, c := range cs[index] {
			if c != 0 {
				v := b ^ num
				ret += c * v
				cs[1-index][v] += c
			}
		}
	}
	return ret
}
