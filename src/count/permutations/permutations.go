package permutations

const mask int64 = 1e9 + 7

func countPermutations(complexity []int) int {
	min := complexity[0]
	n := len(complexity)
	for i := 1; i < n; i++ {
		if complexity[i] <= min {
			return 0
		}
	}
	var out int64 = 1
	i := int64(n - 1)
	for ; i > 1; i-- {
		out = (out * i) % mask
	}
	return int(out)
}
