package parts

func threeEqualParts(arr []int) []int {
	n := 0
	for _, num := range arr {
		if num == 1 {
			n++
		}
	}
	if n%3 != 0 {
		return []int{-1, -1}
	}
	if n == 0 {
		return []int{0, 2}
	}
	part := n / 3
	ret := []int{-1, -1}
	i := 0
	for i < len(arr) && arr[i] == 0 {
		i++
	}
	start := i
	temp := 0
	i = len(arr) - 1
	for temp < part {
		if arr[i] == 1 {
			temp++
		}
		i--
	}
	s1 := i + 1
	n = len(arr) - s1
	for j := 0; j < n; j++ {
		if arr[start+j] != arr[s1+j] {
			return []int{-1, -1}
		}
	}
	ret[0] = start + n - 1
	i = start + n
	for i < len(arr) && arr[i] == 0 {
		i++
	}
	for j := 0; j < n; j++ {
		if arr[start+j] != arr[i+j] {
			return []int{-1, -1}
		}
	}
	ret[1] = i + n
	return ret
}
