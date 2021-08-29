package subarrays

func sumOddLengthSubarrays(arr []int) int {
	prefix_sums := make([]int, len(arr))
	prefix_sums[0] = arr[0]
	for i := 1; i < len(arr); i += 1 {
		prefix_sums[i] = prefix_sums[i-1] + arr[i]
	}

	ret := 0
	for i := 0; i < len(arr); i += 1 {
		for j := i; j < len(arr); j += 2 {
			if i == 0 {
				ret += prefix_sums[j]
			} else {
				ret += prefix_sums[j] - prefix_sums[i-1]
			}
		}
	}
	return ret
}
