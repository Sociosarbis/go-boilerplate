package integer

func findSpecialInteger(arr []int) int {
	n := len(arr)
	t := n / 4
	var count int
	for i, num := range arr {
		count++
		if i+1 < n && arr[i+1] == num {
			continue
		} else {
			if count > t {
				return num
			}
			count = 0
		}
	}
	return arr[0]
}
