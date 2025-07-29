package subarrays

func smallestSubarrays(nums []int) []int {
	indices := [31][]int{}
	for i, num := range nums {
		for j := 0; j < 31; j++ {
			if num&(1<<j) != 0 {
				indices[j] = append(indices[j], i)
			}
		}
	}
	n := len(nums)
	out := make([]int, n)
	curs := [31]int{}
	for i := range nums {
		j := i
		for k := 0; k < 31; k++ {
			for curs[k] < len(indices[k]) && indices[k][curs[k]] < i {
				curs[k]++
			}
			if curs[k] < len(indices[k]) && indices[k][curs[k]] > j {
				j = indices[k][curs[k]]
			}
		}
		out[i] = j - i + 1
	}
	return out
}
