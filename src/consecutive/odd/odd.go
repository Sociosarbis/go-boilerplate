package odd

func threeConsecutiveOdds(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	var count int
	for i := 0; i < 3; i++ {
		if arr[i]&1 == 1 {
			count++
		}
	}
	if count == 3 {
		return true
	}
	for i := 3; i < n; i++ {
		if arr[i-3]&1 == 1 {
			count--
		}
		if arr[i]&1 == 1 {
			count++
		}
		if count == 3 {
			return true
		}
	}
	return false
}
