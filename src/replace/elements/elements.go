package elements

func replaceElements(arr []int) []int {
	n := len(arr)
	temp := arr[n-1]
	arr[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		if arr[i] > temp {
			arr[i], temp = temp, arr[i]
		} else {
			arr[i] = temp
		}
	}
	return arr
}
