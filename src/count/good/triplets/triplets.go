package triplets

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n1 := len(arr)
	n2 := n1 - 1
	n3 := n1 - 2
	var ret int
	for i := 0; i < n3; i++ {
		for j := i + 1; j < n2; j++ {
			if abs(arr[i], arr[j]) <= a {
				for k := j + 1; k < n1; k++ {
					if abs(arr[i], arr[k]) <= c && abs(arr[j], arr[k]) <= b {
						ret++
					}
				}
			}
		}
	}
	return ret
}
